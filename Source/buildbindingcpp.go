/*++

Copyright (C) 2018 Autodesk Inc. (Original Author)

All rights reserved.

Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
this list of conditions and the following disclaimer in the documentation
and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

--*/

//////////////////////////////////////////////////////////////////////////////////////////////////////
// buildbindingcpp.go
// functions to generate C++-bindings of a library's API in form of automatically implemented C++-
// wrapper classes.
//////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"log"
	"path"
	"strings"
	"path/filepath"
)

// BuildBindingCPP builds C++-bindings of a library's API in form of automatically implemented C++-
// wrapper classes.
func BuildBindingCPP(component ComponentDefinition, outputFolder string, outputFolderExample string, indentString string) error {
	namespace := component.NameSpace;
	libraryname := component.LibraryName;
	baseName := component.BaseName;
	forceRecreation := true

	CppHeaderName := path.Join(outputFolder, baseName+".hpp");
	log.Printf("Creating \"%s\"", CppHeaderName)
	hppfile, err :=CreateLanguageFile(CppHeaderName, indentString)
	if err != nil {
		return err
	}

	CppImplName := path.Join(outputFolder, baseName+".cpp");
	log.Printf("Creating \"%s\"", CppImplName)
	cppfile, err :=CreateLanguageFile(CppImplName, indentString)
	if err != nil {
		return err
	}

	WriteLicenseHeader(hppfile.Writer, component,
		fmt.Sprintf("This is an autogenerated C++ Header file in order to allow an easy use\n of %s", libraryname),
		true)
	WriteLicenseHeader(cppfile.Writer, component,
		fmt.Sprintf("This is an autogenerated C++ Wrapper Implementation file in order to allow \nan easy use of %s", libraryname),
		true)

	err = buildCPPHeaderAndImplementation(component, hppfile, cppfile, namespace, baseName)
	if err != nil {
		return err
	}

	if (len(outputFolderExample) > 0) {
		CPPExample := path.Join(outputFolderExample, namespace+"_example"+".cpp");
		if (forceRecreation || !FileExists(CPPExample)) {
			log.Printf("Creating \"%s\"", CPPExample)
			cppexamplefile, err := CreateLanguageFile (CPPExample, "  ")
			if err != nil {
				return err;
			}
			cppexamplefile.WriteCLicenseHeader(component,
				fmt.Sprintf("This is an autogenerated C++ application that demonstrates the\n usage of the C++ bindings of %s", libraryname),
				true)
			buildCppExample(component, cppexamplefile, outputFolder)
		} else {
			log.Printf("Omitting recreation of C++Dynamic example file \"%s\"", CPPExample)
		}

		CPPCMake := path.Join(outputFolderExample, "CMakeLists.txt");
		if (forceRecreation || !FileExists(CPPCMake)) {
			log.Printf("Creating \"%s\"", CPPCMake)
			cppcmake, err := CreateLanguageFile (CPPCMake, "	")
			if err != nil {
				return err;
			}
			cppcmake.WriteCMakeLicenseHeader(component,
				fmt.Sprintf("This is an autogenerated CMake Project that demonstrates the\n usage of the C++ bindings of %s", libraryname),
				true)
			buildCppExampleCMake(component, cppcmake, outputFolder)
		} else {
			log.Printf("Omitting recreation of C++Dynamic example file \"%s\"", CPPCMake)
		}
	}
	return nil
}

func buildCPPHeaderAndImplementation(component ComponentDefinition, w LanguageWriter, cppimplw LanguageWriter, NameSpace string, BaseName string) error {
	// Header start code
	w.Writeln("")
	w.Writeln("#ifndef __%s_CPPHEADER", strings.ToUpper(NameSpace))
	w.Writeln("#define __%s_CPPHEADER", strings.ToUpper(NameSpace))
	w.Writeln("")

	w.Writeln("#include \"%s.h\"", BaseName)

	w.Writeln("#include <string>")
	w.Writeln("#include <memory>")
	w.Writeln("#include <vector>")
	w.Writeln("#include <exception>")
	w.Writeln("")

	w.Writeln("namespace %s {", NameSpace)
	w.Writeln("")

	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Forward Declaration of all classes ")
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("")

	cppClassPrefix := "C" + NameSpace

	w.Writeln("class %sBaseClass;", cppClassPrefix)
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		w.Writeln("class %s%s;", cppClassPrefix, class.ClassName)
	}

	w.Writeln("")

	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Declaration of shared pointer types ")
	w.Writeln("**************************************************************************************************************************/")

	w.Writeln("")

	w.Writeln("typedef std::shared_ptr<%sBaseClass> P%sBaseClass;", cppClassPrefix, NameSpace)
	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		w.Writeln("typedef std::shared_ptr<%s%s> P%s%s;", cppClassPrefix, class.ClassName, NameSpace, class.ClassName)
	}

	w.Writeln("     ")
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Class E%sException ", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("class E%sException : public std::exception {", NameSpace)
	w.Writeln("  protected:")
	w.Writeln("    /**")
	w.Writeln("    * Error code for the Exception.")
	w.Writeln("    */")
	w.Writeln("    %sResult m_errorCode;", NameSpace)
	w.Writeln("    /**")
	w.Writeln("    * Error message for the Exception.")
	w.Writeln("    */")
	w.Writeln("    std::string m_errorMessage;")
	w.Writeln("")
	w.Writeln("  public:")
	w.Writeln("    /**")
	w.Writeln("    * Exception Constructor.")
	w.Writeln("    */")
	w.Writeln("    E%sException (%sResult errorCode, const std::string & sErrorMessage);", NameSpace, NameSpace)
	w.Writeln("")
	w.Writeln("    /**")
	w.Writeln("    * Returns error code")
	w.Writeln("    */")
	w.Writeln("    %sResult getErrorCode () const noexcept;", NameSpace)
	w.Writeln("")
	w.Writeln("    /**")
	w.Writeln("    * Returns error message")
	w.Writeln("    */")
	w.Writeln("    const char* what () const noexcept;")
	w.Writeln("")

	w.Writeln("};")

	w.Writeln("")
	err := writeCPPInputVector(w, NameSpace)
	if err != nil {
		return err
	}
	w.Writeln("")

	w.Writeln("")
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Class %sBaseClass ", cppClassPrefix)
	w.Writeln("**************************************************************************************************************************/")

	w.Writeln("class %sBaseClass {", cppClassPrefix)
	w.Writeln("protected:")

	w.Writeln("  /* Handle to Instance in library*/")
	w.Writeln("  %sHandle m_pHandle;", NameSpace)
	w.Writeln("")
	w.Writeln("  /* Checks for an Error code and raises Exceptions */")
	w.Writeln("  void CheckError(%sResult nResult);", NameSpace)
	w.Writeln("public:")
	w.Writeln("")
	w.Writeln("  /**")
	w.Writeln("  * %sBaseClass::%sBaseClass - Constructor for Base class.", cppClassPrefix, cppClassPrefix)
	w.Writeln("  */")
	w.Writeln("  %sBaseClass(%sHandle pHandle);", cppClassPrefix, NameSpace)
	w.Writeln("")
	w.Writeln("  /**")
	w.Writeln("  * %sBaseClass::~%sBaseClass - Destructor for Base class.", cppClassPrefix, cppClassPrefix)
	w.Writeln("  */")

	w.Writeln("  virtual ~%sBaseClass();", cppClassPrefix)
	w.Writeln("")
	w.Writeln("  /**")
	w.Writeln("  * %sBaseClass::GetHandle - Returns handle to instance.", cppClassPrefix)
	w.Writeln("  */")
	w.Writeln("  %sHandle GetHandle();", NameSpace)
	w.Writeln("};")

	// Implementation start code
	cppimplw.Writeln("#include \"%s.hpp\"", BaseName)
	cppimplw.Writeln("")
	cppimplw.Writeln("#include <vector>")
	cppimplw.Writeln("")
	cppimplw.Writeln("namespace %s {", NameSpace)
	cppimplw.Writeln("")
	cppimplw.Writeln("/*************************************************************************************************************************")
	cppimplw.Writeln(" Class E%sException ", NameSpace)
	cppimplw.Writeln("**************************************************************************************************************************/")
	cppimplw.Writeln("  E%sException::E%sException(%sResult errorCode, const std::string & sErrorMessage)", NameSpace, NameSpace, NameSpace)	
	cppimplw.Writeln("    : m_errorMessage(\"%s Error \" + std::to_string (errorCode) + \" (\" + sErrorMessage + \")\")", NameSpace)
	cppimplw.Writeln("  {")
	cppimplw.Writeln("    m_errorCode = errorCode;")
	cppimplw.Writeln("  }")
	cppimplw.Writeln("")
	cppimplw.Writeln("  %sResult E%sException::getErrorCode () const noexcept", NameSpace, NameSpace)
	cppimplw.Writeln("  {")
	cppimplw.Writeln("    return m_errorCode;")
	cppimplw.Writeln("  }")

	cppimplw.Writeln("")
	cppimplw.Writeln("  const char* E%sException::what () const noexcept", NameSpace)
	cppimplw.Writeln("  {")
	cppimplw.Writeln("    return m_errorMessage.c_str();")
	cppimplw.Writeln("  }")

	cppimplw.Writeln("")
	cppimplw.Writeln("/*************************************************************************************************************************")
	cppimplw.Writeln(" Class %sBaseClass ", cppClassPrefix)
	cppimplw.Writeln("**************************************************************************************************************************/")
	cppimplw.Writeln("")
	cppimplw.Writeln("%sBaseClass::%sBaseClass(%sHandle pHandle)", cppClassPrefix, cppClassPrefix, NameSpace)
	cppimplw.Writeln("{")
	cppimplw.Writeln("  m_pHandle = pHandle;")
	cppimplw.Writeln("}")
	cppimplw.Writeln("")
	cppimplw.Writeln("%sBaseClass::~%sBaseClass()", cppClassPrefix, cppClassPrefix)
	cppimplw.Writeln("{")
	cppimplw.Writeln("  %sWrapper::%s(this);", cppClassPrefix, component.Global.ReleaseMethod)
	cppimplw.Writeln("}")
	cppimplw.Writeln("")
	cppimplw.Writeln("void %sBaseClass::CheckError(%sResult nResult)", cppClassPrefix, NameSpace)
	cppimplw.Writeln("{")
	cppimplw.Writeln("  %sWrapper::CheckError(this, nResult);", cppClassPrefix)
	cppimplw.Writeln("}")
	cppimplw.Writeln("")
	cppimplw.Writeln("%sHandle %sBaseClass::GetHandle()", NameSpace, cppClassPrefix)
	cppimplw.Writeln("{")
	cppimplw.Writeln("  return m_pHandle;")
	cppimplw.Writeln("}")
	cppimplw.Writeln("")

	for i := 0; i < len(component.Classes); i++ {

		class := component.Classes[i]
		cppClassName := cppClassPrefix + class.ClassName

		parentClassName := class.ParentClass
		if parentClassName == "" {
			parentClassName = "BaseClass"
		}
		cppParentClassName := cppClassPrefix + parentClassName

		w.Writeln("     ")
		w.Writeln("/*************************************************************************************************************************")
		w.Writeln(" Class %s ", cppClassName)
		w.Writeln("**************************************************************************************************************************/")
		w.Writeln("class %s : public %s {", cppClassName, cppParentClassName)
		w.Writeln("public:")
		w.Writeln("  ")
		w.Writeln("  /**")
		w.Writeln("  * %s::%s - Constructor for %s class.", cppClassName, cppClassName, class.ClassName)
		w.Writeln("  */")
		w.Writeln("  %s (%sHandle pHandle);", cppClassName, NameSpace)

		cppimplw.Writeln("     ")
		cppimplw.Writeln("/*************************************************************************************************************************")
		cppimplw.Writeln(" Class %s ", cppClassName)
		cppimplw.Writeln("**************************************************************************************************************************/")
		cppimplw.Writeln("/**")
		cppimplw.Writeln("* %s::%s - Constructor for %s class.", cppClassName, cppClassName, class.ClassName)
		cppimplw.Writeln("*/")
		cppimplw.Writeln("%s::%s (%sHandle pHandle)", cppClassName, cppClassName, NameSpace)
		cppimplw.Writeln("  : %s (pHandle)", cppParentClassName)
		cppimplw.Writeln("{ }")

		for j := 0; j < len(class.Methods); j++ {
			method := class.Methods[j]

			err := writeCPPMethod(method, w, cppimplw, NameSpace, class.ClassName, false)
			if err != nil {
				return err
			}

		}

		w.Writeln("};")

	}

	// Global functions
	w.Writeln("     ")
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Class %sWrapper ", cppClassPrefix)
	w.Writeln("**************************************************************************************************************************/")

	w.Writeln("class %sWrapper {", cppClassPrefix)
	w.Writeln("public:")

	w.Writeln("  static void CheckError(%sBaseClass * pBaseClass, %sResult nResult);", cppClassPrefix, NameSpace)

	global := component.Global;
	for j := 0; j < len(global.Methods); j++ {
		method := global.Methods[j]

		err := writeCPPMethod(method, w, cppimplw, NameSpace, "Wrapper", true)
		if err != nil {
			return err
		}
	}

	w.Writeln("};")

	w.Writeln("")
	w.Writeln("};")
	w.Writeln("")
	w.Writeln("#endif // __%s_CPPHEADER", strings.ToUpper(NameSpace))
	w.Writeln("")

	cppimplw.Writeln("")
	cppimplw.Writeln("void %sWrapper::CheckError(%sBaseClass * pBaseClass, %sResult nResult)", cppClassPrefix, cppClassPrefix, NameSpace)
	cppimplw.Writeln("{")	
	cppimplw.Writeln("  if (nResult != 0) {")
	cppimplw.Writeln("    std::string sErrorMessage;")

	if (len (component.Global.ErrorMethod) > 0) {
		cppimplw.Writeln("    if (pBaseClass != nullptr)");
		cppimplw.Writeln("      %s (pBaseClass, sErrorMessage);", component.Global.ErrorMethod);
	}
	cppimplw.Writeln("    throw E%sException (nResult, sErrorMessage);", NameSpace)
	cppimplw.Writeln("  }")
	cppimplw.Writeln("}")
	cppimplw.Writeln("")

	cppimplw.Writeln("")
	cppimplw.Writeln("}; // end namespace %s", NameSpace)
	cppimplw.Writeln("")

	return nil
}

func writeCPPInputVector(w LanguageWriter, NameSpace string) (error) {
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Class C%sInputVector", NameSpace)
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("template <typename T>")
	w.Writeln("class C%sInputVector {", NameSpace)
	w.Writeln("private:")
	w.Writeln("  ")
	w.Writeln("  const T* m_data;")
	w.Writeln("  size_t m_size;")
	w.Writeln("  ")
	w.Writeln("public:")
	w.Writeln("  ")
	w.Writeln("  C%sInputVector( const std::vector<T>& vec)", NameSpace)
	w.Writeln("    : m_data( vec.data() ), m_size( vec.size() )");
	w.Writeln("  {")
	w.Writeln("  }")
	w.Writeln("  ")
	w.Writeln("  C%sInputVector( const T* in_data, size_t in_size)", NameSpace)
	w.Writeln("    : m_data( in_data ), m_size(in_size )");
	w.Writeln("  {")
	w.Writeln("  }")
	w.Writeln("  ")
	w.Writeln("  const T* data() const")
	w.Writeln("  {")
	w.Writeln("    return m_data;")
	w.Writeln("  }")
	w.Writeln("  ")
	w.Writeln("  size_t size() const")
	w.Writeln("  {")
	w.Writeln("    return m_size;")
	w.Writeln("  }")
	w.Writeln("  ")
	w.Writeln("};")
	return nil
}

func getBindingCppParamType (param ComponentDefinitionParam, NameSpace string, isInput bool) (string) {
	cppClassPrefix := "C" + NameSpace;
	switch (param.ParamType) {
		case "uint8":
			return fmt.Sprintf ("%s_uint8", NameSpace);
		case "uint16":
			return fmt.Sprintf ("%s_uint16", NameSpace);
		case "uint32":
			return fmt.Sprintf ("%s_uint32", NameSpace);
		case "uint64":
			return fmt.Sprintf ("%s_uint64", NameSpace);
		case "int8":
			return fmt.Sprintf ("%s_int8", NameSpace);
		case "int16":
			return fmt.Sprintf ("%s_int16", NameSpace);
		case "int32":
			return fmt.Sprintf ("%s_int32", NameSpace);
		case "int64":
			return fmt.Sprintf ("%s_int64", NameSpace);
		case "string":
			return fmt.Sprintf ("std::string");
		case "bool":
			return fmt.Sprintf ("bool");
		case "single":
			return fmt.Sprintf ("float");
		case "basicarray":
			cppBasicType := "";
			switch (param.ParamClass) {
			case "uint8":
				cppBasicType = fmt.Sprintf ("%s_uint8", NameSpace);
			case "uint16":
				cppBasicType = fmt.Sprintf ("%s_uint16", NameSpace);
			case "uint32":
				cppBasicType = fmt.Sprintf ("%s_uint32", NameSpace);
			case "uint64":
				cppBasicType = fmt.Sprintf ("%s_uint64", NameSpace);
			case "int8":
				cppBasicType = fmt.Sprintf ("%s_int8", NameSpace);
			case "int16":
				cppBasicType = fmt.Sprintf ("%s_int16", NameSpace);
			case "int32":
				cppBasicType = fmt.Sprintf ("%s_int32", NameSpace);
			case "int64":
				cppBasicType = fmt.Sprintf ("%s_int64", NameSpace);
			case "bool":
				cppBasicType = "bool";
			case "single":
				cppBasicType =fmt.Sprintf ("%s_single", NameSpace);
			case "double":
				cppBasicType = fmt.Sprintf ("%s_double", NameSpace);
			default:
				log.Fatal ("Invalid parameter type: ", param.ParamClass);
			}
			if (isInput) {
				return fmt.Sprintf ("C%sInputVector<%s>", NameSpace, cppBasicType);
			}
			return fmt.Sprintf ("std::vector<%s>", cppBasicType);
		case "structarray":
			if (isInput) {
				return fmt.Sprintf ("C%sInputVector<s%s%s>", NameSpace, NameSpace, param.ParamClass);
			}
			return fmt.Sprintf ("std::vector<s%s%s>", NameSpace, param.ParamClass);
		case "double":
			return fmt.Sprintf ("%s_double", NameSpace);
		case "enum":
			return fmt.Sprintf ("e%s%s", NameSpace, param.ParamClass);
		case "struct":
			return fmt.Sprintf ("s%s%s", NameSpace, param.ParamClass);
		case "handle":
			if (isInput) {
				return fmt.Sprintf ("%s%s *", cppClassPrefix, param.ParamClass);
			}
			return fmt.Sprintf ("P%s%s", NameSpace, param.ParamClass);
		case "functiontype":
			return fmt.Sprintf ("%s%s", NameSpace, param.ParamClass);
	}
	
	log.Fatal ("Invalid parameter type: ", param.ParamType);
	return "";
}

func getBindingCppVariableName (param ComponentDefinitionParam) (string) {
	switch (param.ParamType) {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
			return "n" + param.ParamName;
		case "string":
			return "s" + param.ParamName;
		case "bool":
			return "b" + param.ParamName;
		case "single":
			return "f" + param.ParamName;
		case "basicarray", "structarray":
			return param.ParamName + "Buffer";
		case "double":
			return "d" + param.ParamName;
		case "enum":
			return "e" + param.ParamName;
		case "struct":
			return param.ParamName;
		case "handle":
			return "p" + param.ParamName;
		case "functiontype":
			return fmt.Sprintf ("p%s", param.ParamName);
	}

	log.Fatal ("Invalid parameter type: ", param.ParamType);
	
	return "";
}

func writeCPPMethod(method ComponentDefinitionMethod, w LanguageWriter, cppimplw LanguageWriter, NameSpace string, ClassName string, isGlobal bool) error {

	CMethodName := ""
	requiresInitCall := false;
	initCallParameters := ""	// usually used to check sizes of buffers
	callParameters := ""
	staticPrefix := ""
	checkErrorCode := ""

	if isGlobal {
		CMethodName = fmt.Sprintf("%s_%s%s", strings.ToLower(NameSpace), strings.ToLower(method.MethodName), method.DLLSuffix)
		staticPrefix = "static "
		checkErrorCode = "CheckError (nullptr,"
	} else {
		CMethodName = fmt.Sprintf("%s_%s_%s%s", strings.ToLower(NameSpace), strings.ToLower(ClassName), strings.ToLower(method.MethodName), method.DLLSuffix)
		callParameters = "m_pHandle"
		initCallParameters = "m_pHandle"
		checkErrorCode = "CheckError ("
	}

	parameters := ""
	returntype := "void"

	definitionCodeLines := []string{}
	functionCodeLines := []string{}
	returnCodeLines := []string{}
	commentcodeLines := []string{}
	postCallCodeLines := []string{}

	cppClassPrefix := "C" + NameSpace
	cppClassName := cppClassPrefix + ClassName

	for k := 0; k < len(method.Params); k++ {

		param := method.Params[k]
		variableName := getBindingCppVariableName(param)

		callParameter := "";
		initCallParameter := "";

		switch param.ParamPass {
		case "in":

			if parameters != "" {
				parameters = parameters + ", "
			}

			cppParamType := getBindingCppParamType(param, NameSpace, true)
			commentcodeLines = append(commentcodeLines, fmt.Sprintf("* @param[in] %s - %s", variableName, param.ParamDescription))

			switch param.ParamType {
			case "string":
				callParameter = variableName + ".c_str()"
				initCallParameter = callParameter;
				parameters = parameters + fmt.Sprintf("const %s & %s", cppParamType, variableName);
			case "struct":
				callParameter = "&" + variableName
				initCallParameter = callParameter;
				parameters = parameters + fmt.Sprintf("const %s & %s", cppParamType, variableName);
			case "structarray", "basicarray":
				callParameter = fmt.Sprintf("(%s_uint64)%s.size(), %s.data()", NameSpace, variableName, variableName);
				initCallParameter = callParameter;
				parameters = parameters + fmt.Sprintf("const %s & %s", cppParamType, variableName);
			case "handle":
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%sHandle h%s = nullptr;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("if (%s != nullptr) {", variableName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("  h%s = %s->GetHandle ();", param.ParamName, variableName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("};"))
				callParameter = "h" + param.ParamName;
				initCallParameter = callParameter;
				parameters = parameters + fmt.Sprintf("%s %s", cppParamType, variableName)

			default:
				callParameter = variableName;
				initCallParameter = callParameter;
				parameters = parameters + fmt.Sprintf("const %s %s", cppParamType, variableName)
			}

		case "out":
			cppParamType := getBindingCppParamType(param, NameSpace, false)
			commentcodeLines = append(commentcodeLines, fmt.Sprintf("* @param[out] %s - %s", variableName, param.ParamDescription))

			if parameters != "" {
				parameters = parameters + ", "
			}
			parameters = parameters + fmt.Sprintf("%s & %s", cppParamType, variableName)

			switch param.ParamType {

			case "string":
				requiresInitCall = true;
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesNeeded%s = 0;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesWritten%s = 0;", NameSpace, param.ParamName))
				initCallParameter = fmt.Sprintf("0, &bytesNeeded%s, nullptr", param.ParamName);
				
				functionCodeLines = append(functionCodeLines, fmt.Sprintf("std::vector<char> buffer%s;", param.ParamName))
				functionCodeLines = append(functionCodeLines, fmt.Sprintf("buffer%s.resize(bytesNeeded%s + 2);", param.ParamName, param.ParamName))

				callParameter = fmt.Sprintf("bytesNeeded%s + 2, &bytesWritten%s, &buffer%s[0]", param.ParamName, param.ParamName, param.ParamName)

				postCallCodeLines = append(postCallCodeLines, fmt.Sprintf("buffer%s[bytesNeeded%s + 1] = 0;", param.ParamName, param.ParamName))
				postCallCodeLines = append(postCallCodeLines, fmt.Sprintf("s%s = std::string(&buffer%s[0]);", param.ParamName, param.ParamName))

			case "handle":
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%sHandle h%s = nullptr;", NameSpace, param.ParamName))
				callParameter = fmt.Sprintf("&h%s", param.ParamName)
				initCallParameter = callParameter;
				postCallCodeLines = append(postCallCodeLines, fmt.Sprintf("p%s = std::make_shared<%s%s> (h%s);", param.ParamName, cppClassPrefix, param.ParamClass, param.ParamName))

			case "structarray", "basicarray":
				requiresInitCall = true;
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint64 elementsNeeded%s = 0;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint64 elementsWritten%s = 0;", NameSpace, param.ParamName))
				initCallParameter = fmt.Sprintf("0, &elementsNeeded%s, nullptr", param.ParamName);

				functionCodeLines = append(functionCodeLines, fmt.Sprintf("%s.resize(elementsNeeded%s);", variableName, param.ParamName))
				callParameter = fmt.Sprintf("elementsNeeded%s, &elementsWritten%s, %s.data()", param.ParamName, param.ParamName, variableName)

			default:
				callParameter = "&" + variableName
				initCallParameter = callParameter
			}

		case "return":

			commentcodeLines = append(commentcodeLines, fmt.Sprintf("* @return %s", param.ParamDescription))
			returntype = getBindingCppParamType(param, NameSpace, false)

			switch param.ParamType {
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "bool", "single", "double":
				callParameter = fmt.Sprintf("&result%s", param.ParamName)
				initCallParameter = callParameter;
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s result%s = 0;", returntype, param.ParamName))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return result%s;", param.ParamName))

			case "string":
				requiresInitCall = true;
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesNeeded%s = 0;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesWritten%s = 0;", NameSpace, param.ParamName))
				initCallParameter = fmt.Sprintf("0, &bytesNeeded%s, nullptr", param.ParamName);

				functionCodeLines = append(functionCodeLines, fmt.Sprintf("std::vector<char> buffer%s;", param.ParamName))
				functionCodeLines = append(functionCodeLines, fmt.Sprintf("buffer%s.resize(bytesNeeded%s + 2);", param.ParamName, param.ParamName))

				callParameter = fmt.Sprintf("bytesNeeded%s + 2, &bytesWritten%s, &buffer%s[0]", param.ParamName, param.ParamName, param.ParamName)

				returnCodeLines = append(returnCodeLines, fmt.Sprintf("buffer%s[bytesNeeded%s + 1] = 0;", param.ParamName, param.ParamName))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return std::string(&buffer%s[0]);", param.ParamName))

			case "enum":
				callParameter = fmt.Sprintf("&result%s", param.ParamName)
				initCallParameter = callParameter;
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("e%s%s result%s = (e%s%s) 0;", NameSpace, param.ParamClass, param.ParamName, NameSpace, param.ParamClass))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return result%s;", param.ParamName))

			case "struct":
				callParameter = fmt.Sprintf("&result%s", param.ParamName)
				initCallParameter = callParameter;
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("s%s%s result%s;", NameSpace, param.ParamClass, param.ParamName))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return result%s;", param.ParamName))

			case "handle":
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%sHandle h%s = nullptr;", NameSpace, param.ParamName))
				callParameter = fmt.Sprintf("&h%s", param.ParamName)
				initCallParameter = callParameter;
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return std::make_shared<%s%s> (h%s);", cppClassPrefix, param.ParamClass, param.ParamName))

			case "basicarray":
				return fmt.Errorf("can not return basicarray \"%s\" for %s.%s (%s)", param.ParamPass, ClassName, method.MethodName, param.ParamName)

			case "structarray":
				return fmt.Errorf("can not return structarray \"%s\" for %s.%s (%s)", param.ParamPass, ClassName, method.MethodName, param.ParamName)

			default:
				return fmt.Errorf("invalid method parameter type \"%s\" for %s.%s (%s)", param.ParamType, ClassName, method.MethodName, param.ParamName)
			}

		default:
			return fmt.Errorf("invalid method parameter passing \"%s\" for %s.%s (%s)", param.ParamPass, ClassName, method.MethodName, param.ParamName)
		}

		if callParameters != "" {
			callParameters = callParameters + ", "
		}
		callParameters = callParameters + callParameter;
		if (initCallParameters != "") {
			initCallParameters = initCallParameters + ", ";
		}
		initCallParameters = initCallParameters + initCallParameter;

	}

	w.Writeln("")
	w.Writeln("  /**")
	w.Writeln("  * %s::%s - %s", cppClassName, method.MethodName, method.MethodDescription)
	w.Writelns("  ", commentcodeLines)
	w.Writeln("  */")
	w.Writeln("  %s%s %s (%s);", staticPrefix, returntype, method.MethodName, parameters)

	cppimplw.Writeln("")
	cppimplw.Writeln("/**")
	cppimplw.Writeln("* %s::%s - %s", cppClassName, method.MethodName, method.MethodDescription)
	cppimplw.Writelns("", commentcodeLines)
	cppimplw.Writeln("*/")
	cppimplw.Writeln("%s %s::%s (%s)", returntype, cppClassName, method.MethodName, parameters)
	cppimplw.Writeln("{")
	cppimplw.Writelns("  ", definitionCodeLines)
	if (requiresInitCall) {
		cppimplw.Writeln("  %s %s (%s) );", checkErrorCode, CMethodName, initCallParameters)
	}
	cppimplw.Writelns("  ", functionCodeLines)
	cppimplw.Writeln("  %s %s (%s) );", checkErrorCode, CMethodName, callParameters)
	cppimplw.Writelns("  ", postCallCodeLines)
	cppimplw.Writelns("  ", returnCodeLines)
	cppimplw.Writeln("}")
	

	return nil
}


func buildCppExample(componentdefinition ComponentDefinition, w LanguageWriter, outputFolder string) error {
	NameSpace := componentdefinition.NameSpace
	BaseName := componentdefinition.BaseName

	w.Writeln("#include <iostream>")
	w.Writeln("#include \"%s.hpp\"", strings.ToLower(BaseName) )
	w.Writeln("")
	w.Writeln("")

	w.Writeln("int main()")
	w.Writeln("{")
	w.Writeln("  try")
	w.Writeln("  {")
	w.Writeln("    unsigned int nMajor, nMinor, nMicro;")
	w.Writeln("    %s::C%sWrapper::%s(nMajor, nMinor, nMicro);", NameSpace, NameSpace, componentdefinition.Global.VersionMethod)
	w.Writeln("    std::cout << \"%s.Version = \" << nMajor << \".\" << nMinor << \".\" << nMicro << std::endl;", NameSpace);
	w.Writeln("  }")
	w.Writeln("  catch (std::exception &e)")
	w.Writeln("  {")
	w.Writeln("    std::cout << e.what() << std::endl;")
	w.Writeln("    return 1;")
	w.Writeln("  }")
	w.Writeln("  return 0;")
	w.Writeln("}")
	w.Writeln("")

	return nil
}

func buildCppExampleCMake(componentdefinition ComponentDefinition, w LanguageWriter, outputFolder string) error {
	NameSpace := componentdefinition.NameSpace
	BaseName := componentdefinition.BaseName

	w.Writeln("cmake_minimum_required(VERSION 3.5)")

	w.Writeln("")
	w.Writeln("project(%sExample_CPP)", NameSpace)
	w.Writeln("set (CMAKE_CXX_STANDARD 11)")
	// TODO: calculate relative path from ExampleOutputFolder to OUTPUTFOLDER based on CURRENT_SOURCE_DIR
	linkfolder := strings.Replace(outputFolder, string(filepath.Separator), "/", -2)
	w.Writeln("link_directories(\"%s\") # TODO: put the correct path of the import library here", linkfolder)

	outputFolder = strings.Replace(outputFolder, string(filepath.Separator), "/", -1)
	w.Writeln("add_executable(%sExample_CPP \"${CMAKE_CURRENT_SOURCE_DIR}/%s_example.cpp\"", NameSpace, NameSpace)
	w.Writeln("  \"%s/%s.cpp\")", outputFolder, BaseName)
	w.Writeln("target_link_libraries(%sExample_CPP %s)", NameSpace, BaseName)

	// TODO: calculate relative path from ExampleOutputFolder to OUTPUTFOLDER based on CURRENT_SOURCE_DIR
	w.Writeln("target_include_directories(%sExample_CPP PRIVATE \"%s\")", NameSpace, outputFolder)
	
	return nil
}
