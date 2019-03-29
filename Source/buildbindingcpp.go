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
	"path/filepath"
	"strings"
)


func writeCppBaseClassMethods(component ComponentDefinition, baseClass ComponentDefinitionClass, w LanguageWriter, NameSpace string, BaseName string, cppClassPrefix string) {
	cppBaseClassName := cppClassPrefix + baseClass.ClassName
	w.Writeln("protected:")
	w.Writeln("  /* Handle to Instance in library*/")
	w.Writeln("  %sHandle m_pHandle;", NameSpace)
	w.Writeln("")
	w.Writeln("  /* Checks for an Error code and raises Exceptions */")
	w.Writeln("  void CheckError(%sResult nResult);", NameSpace)
	w.Writeln("")
	w.Writeln("  /**")
	w.Writeln("  * %s::%s - Constructor for Base class.", cppBaseClassName, cppBaseClassName)
	w.Writeln("  */")
	w.Writeln("  %s(%sHandle pHandle);", cppBaseClassName, NameSpace)
	w.Writeln("")
	w.Writeln("  /**")
	w.Writeln("  * %s::~%s - Destructor for Base class.", cppBaseClassName, cppBaseClassName)
	w.Writeln("  */")
	w.Writeln("  virtual ~%s();", cppBaseClassName)
	w.Writeln("")
	w.Writeln("public:")
	w.Writeln("  /**")
	w.Writeln("  * %s::GetHandle - Returns handle to instance.", cppBaseClassName)
	w.Writeln("  */")
	w.Writeln("  %sHandle GetHandle();", NameSpace)
	w.Writeln("  ")
}

func writeCppBaseClassDefinitions(component ComponentDefinition, baseClass ComponentDefinitionClass, w LanguageWriter, NameSpace string, BaseName string, cppClassPrefix string) error {
	cppBaseClassName := cppClassPrefix + baseClass.ClassName

	w.Writeln("void %s::CheckError(%sResult nResult)", cppBaseClassName, NameSpace)
	w.Writeln("{")
	w.Writeln("  %sWrapper::CheckError(this, nResult);", cppClassPrefix)
	w.Writeln("}")
	w.Writeln("")
	w.Writeln("%s::%s(%sHandle pHandle)", cppBaseClassName, cppBaseClassName, NameSpace)
	w.Writeln("  : m_pHandle (pHandle)")
	w.Writeln("{")
	w.Writeln("}")
	w.Writeln("")
	w.Writeln("%s::~%s()", cppBaseClassName, cppBaseClassName)
	w.Writeln("{")
	w.Writeln("  %sWrapper::%s(this);", cppClassPrefix, component.Global.ReleaseMethod)
	w.Writeln("}")
	w.Writeln("")

	w.Writeln("%sHandle %s::GetHandle()", NameSpace, cppBaseClassName)
	w.Writeln("{")
	w.Writeln("  return m_pHandle;")
	w.Writeln("}")
	w.Writeln("")
	return nil
}

func buildCPPHeaderAndImplementation(component ComponentDefinition, w LanguageWriter, cppimplw LanguageWriter, NameSpace string, BaseName string) error {
	cppClassPrefix := "C"
	cppBaseClassName := cppClassPrefix + component.Global.BaseClassName

	// Header start code
	w.Writeln("")
	w.Writeln("#ifndef __%s_CPPHEADER", strings.ToUpper(NameSpace))
	w.Writeln("#define __%s_CPPHEADER", strings.ToUpper(NameSpace))
	w.Writeln("")

	w.Writeln("#include \"%s_abi.hpp\"", BaseName)

	w.Writeln("#include <string>")
	w.Writeln("#include <memory>")
	w.Writeln("#include <vector>")
	w.Writeln("#include <exception>")
	w.Writeln("")

	w.Writeln("namespace %s {", NameSpace)
	w.Writeln("")

	buildBindingCPPAllForwardDeclarations(component, w, NameSpace, cppClassPrefix)

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
	w.Writeln("    E%sException(%sResult errorCode, const std::string & sErrorMessage);", NameSpace, NameSpace)
	w.Writeln("")
	w.Writeln("    /**")
	w.Writeln("    * Returns error code")
	w.Writeln("    */")
	w.Writeln("    %sResult getErrorCode() const noexcept;", NameSpace)
	w.Writeln("")
	w.Writeln("    /**")
	w.Writeln("    * Returns error message")
	w.Writeln("    */")
	w.Writeln("    const char* what() const noexcept;")
	w.Writeln("")

	w.Writeln("};")

	w.Writeln("")
	err := writeCPPInputVector(w, NameSpace)
	if err != nil {
		return err
	}
	w.Writeln("")

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

	for i := 0; i < len(component.Classes); i++ {
		class := component.Classes[i]
		cppClassName := cppClassPrefix + class.ClassName

		cppParentClassName := ""
		inheritanceSpecifier := ""
		if !component.isBaseClass(class) {
			if class.ParentClass == "" {
				cppParentClassName = cppClassPrefix + component.Global.BaseClassName
			} else {
				cppParentClassName = cppClassPrefix + class.ParentClass
			}
			inheritanceSpecifier = fmt.Sprintf(": public %s ", cppParentClassName)
		}

		w.Writeln("     ")
		w.Writeln("/*************************************************************************************************************************")
		w.Writeln(" Class %s ", cppClassName)
		w.Writeln("**************************************************************************************************************************/")
		w.Writeln("class %s %s {", cppClassName, inheritanceSpecifier)

		if component.isBaseClass(class) {
			writeCppBaseClassMethods(component, class, w, NameSpace, BaseName, cppClassPrefix)
		} else {
			w.Writeln("public:")
			w.Writeln("  /**")
			w.Writeln("  * %s::%s - Constructor for %s class.", cppClassName, cppClassName, class.ClassName)
			w.Writeln("  */")
			w.Writeln("  %s (%sHandle pHandle);", cppClassName, NameSpace)
		}

		cppimplw.Writeln("     ")
		cppimplw.Writeln("/*************************************************************************************************************************")
		cppimplw.Writeln(" Class %s ", cppClassName)
		cppimplw.Writeln("**************************************************************************************************************************/")
		cppimplw.Writeln("/**")
		cppimplw.Writeln("* %s::%s - Constructor for %s class.", cppClassName, cppClassName, class.ClassName)
		cppimplw.Writeln("*/")
		if !component.isBaseClass(class) {
			cppimplw.Writeln("%s::%s (%sHandle pHandle)", cppClassName, cppClassName, NameSpace)
			cppimplw.Writeln("  : %s (pHandle)", cppParentClassName)
			cppimplw.Writeln("{ }")
		} else {
			err = writeCppBaseClassDefinitions(component, class, cppimplw, NameSpace, BaseName, cppClassPrefix)
			if err != nil {
				return err
			}
		}

		if len(class.Methods) > 0 {
			w.Writeln("public:")
			w.Writeln("  ")
		}
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
	w.Writeln("  static void CheckError(%s * pBaseClass, %sResult nResult);", cppBaseClassName, NameSpace)

	global := component.Global
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
	cppimplw.Writeln("void %sWrapper::CheckError(%s * pBaseClass, %sResult nResult)", cppClassPrefix, cppBaseClassName, NameSpace)
	cppimplw.Writeln("{")
	cppimplw.Writeln("  if (nResult != 0) {")
	cppimplw.Writeln("    std::string sErrorMessage;")

	cppimplw.Writeln("    if (pBaseClass != nullptr)")
	cppimplw.Writeln("      %s(pBaseClass, sErrorMessage);", component.Global.ErrorMethod)

	cppimplw.Writeln("    throw E%sException (nResult, sErrorMessage);", NameSpace)
	cppimplw.Writeln("  }")
	cppimplw.Writeln("}")
	cppimplw.Writeln("")

	cppimplw.Writeln("")
	cppimplw.Writeln("}; // end namespace %s", NameSpace)
	cppimplw.Writeln("")

	return nil
}

func writeCPPInputVector(w LanguageWriter, NameSpace string) error {
	w.Writeln("/*************************************************************************************************************************")
	w.Writeln(" Class CInputVector")
	w.Writeln("**************************************************************************************************************************/")
	w.Writeln("template <typename T>")
	w.Writeln("class CInputVector {")
	w.Writeln("private:")
	w.Writeln("  ")
	w.Writeln("  const T* m_data;")
	w.Writeln("  size_t m_size;")
	w.Writeln("  ")
	w.Writeln("public:")
	w.Writeln("  ")
	w.Writeln("  CInputVector( const std::vector<T>& vec)")
	w.Writeln("    : m_data( vec.data() ), m_size( vec.size() )")
	w.Writeln("  {")
	w.Writeln("  }")
	w.Writeln("  ")
	w.Writeln("  CInputVector( const T* in_data, size_t in_size)")
	w.Writeln("    : m_data( in_data ), m_size(in_size )")
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
	w.Writeln("")
	w.Writeln("// declare deprecated class name")
	w.Writeln("template<typename T>")
	w.Writeln("using C%sInputVector = CInputVector<T>;", NameSpace)
	return nil
}

func getBindingCppParamType(paramType string, paramClass string, NameSpace string, isInput bool) string {
	cppClassPrefix := "C"
	switch paramType {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double":
		return fmt.Sprintf("%s_%s", NameSpace, paramType)
	case "string":
		return fmt.Sprintf("std::string")
	case "bool":
		return fmt.Sprintf("bool")
	case "pointer":
		return fmt.Sprintf("%s_pvoid", NameSpace)
	case "basicarray":
		cppBasicType := ""
		switch paramClass {
		case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "single", "double",
			"bool", "pointer":
			cppBasicType = getBindingCppParamType(paramClass, "", NameSpace, isInput)
		default:
			log.Fatal("Invalid parameter type: ", paramClass)
		}
		if isInput {
			return fmt.Sprintf("CInputVector<%s>", cppBasicType)
		}
		return fmt.Sprintf("std::vector<%s>", cppBasicType)
	case "structarray":
		if isInput {
			return fmt.Sprintf("CInputVector<s%s>", paramClass)
		}
		return fmt.Sprintf("std::vector<s%s>", paramClass)
	case "enum":
		return fmt.Sprintf("e%s", paramClass)
	case "struct":
		return fmt.Sprintf("s%s", paramClass)
	case "class":
		if isInput {
			return fmt.Sprintf("%s%s *", cppClassPrefix, paramClass)
		}
		return fmt.Sprintf("P%s", paramClass)
	case "functiontype":
		return fmt.Sprintf("%s", paramClass)
	}
	log.Fatal("Invalid parameter type: ", paramType)
	return ""
}

func getBindingCppVariableName(param ComponentDefinitionParam) string {
	switch param.ParamType {
	case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64":
		return "n" + param.ParamName
	case "string":
		return "s" + param.ParamName
	case "bool":
		return "b" + param.ParamName
	case "single":
		return "f" + param.ParamName
	case "basicarray", "structarray":
		return param.ParamName + "Buffer"
	case "double":
		return "d" + param.ParamName
	case "pointer":
		return "p" + param.ParamName
	case "enum":
		return "e" + param.ParamName
	case "struct":
		return param.ParamName
	case "class":
		return "p" + param.ParamName
	case "functiontype":
		return fmt.Sprintf("p%s", param.ParamName)
	}

	log.Fatal("Invalid parameter type: ", param.ParamType)

	return ""
}

func writeCPPMethod(method ComponentDefinitionMethod, w LanguageWriter, cppimplw LanguageWriter, NameSpace string, ClassName string, isGlobal bool) error {

	CMethodName := ""
	requiresInitCall := false
	initCallParameters := "" // usually used to check sizes of buffers
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

	cppClassPrefix := "C"
	cppClassName := cppClassPrefix + ClassName

	for k := 0; k < len(method.Params); k++ {

		param := method.Params[k]
		variableName := getBindingCppVariableName(param)

		callParameter := ""
		initCallParameter := ""

		switch param.ParamPass {
		case "in":

			if parameters != "" {
				parameters = parameters + ", "
			}

			cppParamType := getBindingCppParamType(param.ParamType, param.ParamClass, NameSpace, true)
			commentcodeLines = append(commentcodeLines, fmt.Sprintf("* @param[in] %s - %s", variableName, param.ParamDescription))

			switch param.ParamType {
			case "string":
				callParameter = variableName + ".c_str()"
				initCallParameter = callParameter
				parameters = parameters + fmt.Sprintf("const %s & %s", cppParamType, variableName)
			case "struct":
				callParameter = "&" + variableName
				initCallParameter = callParameter
				parameters = parameters + fmt.Sprintf("const %s & %s", cppParamType, variableName)
			case "structarray", "basicarray":
				callParameter = fmt.Sprintf("(%s_uint64)%s.size(), %s.data()", NameSpace, variableName, variableName)
				initCallParameter = callParameter
				parameters = parameters + fmt.Sprintf("const %s & %s", cppParamType, variableName)
			case "class":
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%sHandle h%s = nullptr;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("if (%s != nullptr) {", variableName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("  h%s = %s->GetHandle ();", param.ParamName, variableName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("};"))
				callParameter = "h" + param.ParamName
				initCallParameter = callParameter
				parameters = parameters + fmt.Sprintf("%s %s", cppParamType, variableName)

			default:
				callParameter = variableName
				initCallParameter = callParameter
				parameters = parameters + fmt.Sprintf("const %s %s", cppParamType, variableName)
			}

		case "out":
			cppParamType := getBindingCppParamType(param.ParamType, param.ParamClass, NameSpace, false)
			commentcodeLines = append(commentcodeLines, fmt.Sprintf("* @param[out] %s - %s", variableName, param.ParamDescription))

			if parameters != "" {
				parameters = parameters + ", "
			}
			parameters = parameters + fmt.Sprintf("%s & %s", cppParamType, variableName)

			switch param.ParamType {

			case "string":
				requiresInitCall = true
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesNeeded%s = 0;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesWritten%s = 0;", NameSpace, param.ParamName))
				initCallParameter = fmt.Sprintf("0, &bytesNeeded%s, nullptr", param.ParamName)

				functionCodeLines = append(functionCodeLines, fmt.Sprintf("std::vector<char> buffer%s;", param.ParamName))
				functionCodeLines = append(functionCodeLines, fmt.Sprintf("buffer%s.resize(bytesNeeded%s + 2);", param.ParamName, param.ParamName))

				callParameter = fmt.Sprintf("bytesNeeded%s + 2, &bytesWritten%s, &buffer%s[0]", param.ParamName, param.ParamName, param.ParamName)

				postCallCodeLines = append(postCallCodeLines, fmt.Sprintf("buffer%s[bytesNeeded%s + 1] = 0;", param.ParamName, param.ParamName))
				postCallCodeLines = append(postCallCodeLines, fmt.Sprintf("s%s = std::string(&buffer%s[0]);", param.ParamName, param.ParamName))

			case "class":
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%sHandle h%s = nullptr;", NameSpace, param.ParamName))
				callParameter = fmt.Sprintf("&h%s", param.ParamName)
				initCallParameter = callParameter
				postCallCodeLines = append(postCallCodeLines, fmt.Sprintf("p%s = std::make_shared<%s%s> (h%s);", param.ParamName, cppClassPrefix, param.ParamClass, param.ParamName))

			case "structarray", "basicarray":
				requiresInitCall = true
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint64 elementsNeeded%s = 0;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint64 elementsWritten%s = 0;", NameSpace, param.ParamName))
				initCallParameter = fmt.Sprintf("0, &elementsNeeded%s, nullptr", param.ParamName)

				functionCodeLines = append(functionCodeLines, fmt.Sprintf("%s.resize(elementsNeeded%s);", variableName, param.ParamName))
				callParameter = fmt.Sprintf("elementsNeeded%s, &elementsWritten%s, %s.data()", param.ParamName, param.ParamName, variableName)

			default:
				callParameter = "&" + variableName
				initCallParameter = callParameter
			}

		case "return":
			commentcodeLines = append(commentcodeLines, fmt.Sprintf("* @return %s", param.ParamDescription))
			returntype = getBindingCppParamType(param.ParamType, param.ParamClass, NameSpace, false)

			switch param.ParamType {
			case "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "bool", "single", "double", "pointer":
				callParameter = fmt.Sprintf("&result%s", param.ParamName)
				initCallParameter = callParameter
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s result%s = 0;", returntype, param.ParamName))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return result%s;", param.ParamName))

			case "string":
				requiresInitCall = true
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesNeeded%s = 0;", NameSpace, param.ParamName))
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%s_uint32 bytesWritten%s = 0;", NameSpace, param.ParamName))
				initCallParameter = fmt.Sprintf("0, &bytesNeeded%s, nullptr", param.ParamName)

				functionCodeLines = append(functionCodeLines, fmt.Sprintf("std::vector<char> buffer%s;", param.ParamName))
				functionCodeLines = append(functionCodeLines, fmt.Sprintf("buffer%s.resize(bytesNeeded%s + 2);", param.ParamName, param.ParamName))

				callParameter = fmt.Sprintf("bytesNeeded%s + 2, &bytesWritten%s, &buffer%s[0]", param.ParamName, param.ParamName, param.ParamName)

				returnCodeLines = append(returnCodeLines, fmt.Sprintf("buffer%s[bytesNeeded%s + 1] = 0;", param.ParamName, param.ParamName))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return std::string(&buffer%s[0]);", param.ParamName))

			case "enum":
				callParameter = fmt.Sprintf("&result%s", param.ParamName)
				initCallParameter = callParameter
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("e%s result%s = (e%s) 0;", param.ParamClass, param.ParamName, param.ParamClass))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return result%s;", param.ParamName))

			case "struct":
				callParameter = fmt.Sprintf("&result%s", param.ParamName)
				initCallParameter = callParameter
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("s%s result%s;", param.ParamClass, param.ParamName))
				returnCodeLines = append(returnCodeLines, fmt.Sprintf("return result%s;", param.ParamName))

			case "class":
				definitionCodeLines = append(definitionCodeLines, fmt.Sprintf("%sHandle h%s = nullptr;", NameSpace, param.ParamName))
				callParameter = fmt.Sprintf("&h%s", param.ParamName)
				initCallParameter = callParameter
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
		callParameters = callParameters + callParameter
		if initCallParameters != "" {
			initCallParameters = initCallParameters + ", "
		}
		initCallParameters = initCallParameters + initCallParameter

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
	if requiresInitCall {
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
	w.Writeln("#include \"%s.hpp\"", strings.ToLower(BaseName))
	w.Writeln("")
	w.Writeln("")

	w.Writeln("int main()")
	w.Writeln("{")
	w.Writeln("  try")
	w.Writeln("  {")
	w.Writeln("    unsigned int nMajor, nMinor, nMicro;")
	w.Writeln("    std::string sPreReleaseInfo, sBuildInfo;")
	w.Writeln("    %s::CWrapper::%s(nMajor, nMinor, nMicro, sPreReleaseInfo, sBuildInfo);", NameSpace, componentdefinition.Global.VersionMethod)
	w.Writeln("    std::cout << \"%s.Version = \" << nMajor << \".\" << nMinor << \".\" << nMicro;", NameSpace)
	w.Writeln("    if (!sPreReleaseInfo.empty())")
	w.Writeln("      std::cout << \"-\" << sPreReleaseInfo;")
	w.Writeln("    if (!sBuildInfo.empty())")
	w.Writeln("      std::cout << \"+\" << sBuildInfo;")
	w.Writeln("    std::cout << std::endl;")

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
