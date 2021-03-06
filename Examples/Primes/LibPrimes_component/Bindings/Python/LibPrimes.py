'''++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0.

Abstract: This is an autogenerated Python file in order to allow an easy
 use of Prime Numbers Library

Interface version: 1.2.0

'''


import ctypes
import platform
import enum

name = "libprimes"

'''Definition of domain specific exception
'''
class ELibPrimesException(Exception):
	def __init__(self, code, message = ''):
		self._code = code
		self._message = message
	
	def __str__(self):
		if self._message:
			return 'LibPrimesException ' + str(self._code) + ': '+ str(self._message)
		return 'LibPrimesException ' + str(self._code)

'''Definition of binding API version
'''
class BindingVersion(enum.IntEnum):
	MAJOR = 1
	MINOR = 2
	MICRO = 0

'''Definition Error Codes
'''
class ErrorCodes(enum.IntEnum):
	SUCCESS = 0
	NOTIMPLEMENTED = 1
	INVALIDPARAM = 2
	INVALIDCAST = 3
	BUFFERTOOSMALL = 4
	GENERICEXCEPTION = 5
	COULDNOTLOADLIBRARY = 6
	COULDNOTFINDLIBRARYEXPORT = 7
	INCOMPATIBLEBINARYVERSION = 8
	NORESULTAVAILABLE = 9
	CALCULATIONABORTED = 10

'''Definition of Structs
'''
'''Definition of PrimeFactor
'''
class PrimeFactor(ctypes.Structure):
	_pack_ = 1
	_fields_ = [
		("Prime", ctypes.c_uint64), 
		("Multiplicity", ctypes.c_uint32)
	]

'''Definition of Function Types
'''
'''Definition of ProgressCallback
		Callback to report calculation progress and query whether it should be aborted
'''
ProgressCallback = ctypes.CFUNCTYPE(ctypes.c_void_p, ctypes.c_float, ctypes.POINTER(ctypes.c_bool))


'''Wrapper Class Implementation
'''
class Wrapper:

	def __init__(self, libraryName = None):
		ending = ''
		if platform.system() == 'Windows':
			ending = 'dll'
		elif platform.system() == 'Linux':
			ending = 'so'
		elif platform.system() == 'Darwin':
			ending = 'dylib'
		else:
			raise ELibPrimesException(ErrorCodes.COULDNOTLOADLIBRARY)
		
		if (not libraryName):
			libraryName = os.path.join(os.path.dirname(os.path.realpath(__file__)),'libprimes')
		path = libraryName + '.' + ending
		
		try:
			self.lib = ctypes.CDLL(path)
		except Exception as e:
			raise ELibPrimesException(ErrorCodes.COULDNOTLOADLIBRARY, str(e) + '| "'+path + '"' )
		
		self._loadFunctionTable()
		
		self._checkBinaryVersion()
	
	def _loadFunctionTable(self):
		try:
			self.lib.libprimes_getversion.restype = ctypes.c_int64
			self.lib.libprimes_getversion.argtypes = [ctypes.POINTER(ctypes.c_uint32), ctypes.POINTER(ctypes.c_uint32), ctypes.POINTER(ctypes.c_uint32)]
			
			self.lib.libprimes_getlasterror.restype = ctypes.c_int64
			self.lib.libprimes_getlasterror.argtypes = [ctypes.c_void_p, ctypes.c_uint64, ctypes.POINTER(ctypes.c_uint64), ctypes.c_char_p, ctypes.POINTER(ctypes.c_bool)]
			
			self.lib.libprimes_releaseinstance.restype = ctypes.c_int64
			self.lib.libprimes_releaseinstance.argtypes = [ctypes.c_void_p]
			
			self.lib.libprimes_createfactorizationcalculator.restype = ctypes.c_int64
			self.lib.libprimes_createfactorizationcalculator.argtypes = [ctypes.POINTER(ctypes.c_void_p)]
			
			self.lib.libprimes_createsievecalculator.restype = ctypes.c_int64
			self.lib.libprimes_createsievecalculator.argtypes = [ctypes.POINTER(ctypes.c_void_p)]
			
			self.lib.libprimes_setjournal.restype = ctypes.c_int64
			self.lib.libprimes_setjournal.argtypes = [ctypes.c_char_p]
			
			self.lib.libprimes_calculator_getvalue.restype = ctypes.c_int64
			self.lib.libprimes_calculator_getvalue.argtypes = [ctypes.c_void_p, ctypes.POINTER(ctypes.c_uint64)]
			
			self.lib.libprimes_calculator_setvalue.restype = ctypes.c_int64
			self.lib.libprimes_calculator_setvalue.argtypes = [ctypes.c_void_p, ctypes.c_uint64]
			
			self.lib.libprimes_calculator_calculate.restype = ctypes.c_int64
			self.lib.libprimes_calculator_calculate.argtypes = [ctypes.c_void_p]
			
			self.lib.libprimes_calculator_setprogresscallback.restype = ctypes.c_int64
			self.lib.libprimes_calculator_setprogresscallback.argtypes = [ctypes.c_void_p, ProgressCallback]
			
			self.lib.libprimes_factorizationcalculator_getprimefactors.restype = ctypes.c_int64
			self.lib.libprimes_factorizationcalculator_getprimefactors.argtypes = [ctypes.c_void_p, ctypes.c_uint64, ctypes.POINTER(ctypes.c_uint64), ctypes.POINTER(PrimeFactor)]
			
			self.lib.libprimes_sievecalculator_getprimes.restype = ctypes.c_int64
			self.lib.libprimes_sievecalculator_getprimes.argtypes = [ctypes.c_void_p, ctypes.c_uint64, ctypes.POINTER(ctypes.c_uint64), ctypes.POINTER(ctypes.c_uint64)]
			
		except AttributeError as ae:
			raise ELibPrimesException(ErrorCodes.COULDNOTFINDLIBRARYEXPORT, ae.args[0])
	
	def _checkBinaryVersion(self):
		nMajor, nMinor, _ = self.GetVersion()
		if (nMajor != BindingVersion.MAJOR) or (nMinor < BindingVersion.MINOR):
			raise ELibPrimesException(ErrorCodes.INCOMPATIBLEBINARYVERSION)
	
	def checkError(self, instance, errorCode):
		if errorCode != ErrorCodes.SUCCESS.value:
			if instance:
				if instance._wrapper != self:
					raise ELibPrimesException(ErrorCodes.INVALIDCAST, 'invalid wrapper call')
			message,_ = self.GetLastError(instance)
			raise ELibPrimesException(errorCode, message)
	
	def GetVersion(self):
		pMajor = ctypes.c_uint32()
		pMinor = ctypes.c_uint32()
		pMicro = ctypes.c_uint32()
		self.checkError(None, self.lib.libprimes_getversion(pMajor, pMinor, pMicro))
		return pMajor.value, pMinor.value, pMicro.value
	
	def GetLastError(self, InstanceObject):
		nErrorMessageBufferSize = ctypes.c_uint64(0)
		nErrorMessageNeededChars = ctypes.c_uint64(0)
		pErrorMessageBuffer = ctypes.c_char_p(None)
		pHasError = ctypes.c_bool()
		self.checkError(None, self.lib.libprimes_getlasterror(InstanceObject._handle, nErrorMessageBufferSize, nErrorMessageNeededChars, pErrorMessageBuffer, pHasError))
		nErrorMessageBufferSize = ctypes.c_uint64(nErrorMessageNeededChars.value + 2)
		pErrorMessageBuffer = (ctypes.c_char * (nErrorMessageNeededChars.value + 2))()
		self.checkError(None, self.lib.libprimes_getlasterror(InstanceObject._handle, nErrorMessageBufferSize, nErrorMessageNeededChars, pErrorMessageBuffer, pHasError))
		return pErrorMessageBuffer.value.decode(), pHasError.value
	
	def ReleaseInstance(self, InstanceObject):
		self.checkError(None, self.lib.libprimes_releaseinstance(InstanceObject._handle))
	
	def CreateFactorizationCalculator(self):
		InstanceHandle = ctypes.c_void_p()
		self.checkError(None, self.lib.libprimes_createfactorizationcalculator(InstanceHandle))
		InstanceObject = FactorizationCalculator(InstanceHandle, self)
		return InstanceObject
	
	def CreateSieveCalculator(self):
		InstanceHandle = ctypes.c_void_p()
		self.checkError(None, self.lib.libprimes_createsievecalculator(InstanceHandle))
		InstanceObject = SieveCalculator(InstanceHandle, self)
		return InstanceObject
	
	def SetJournal(self, FileName):
		pFileName = ctypes.c_char_p(str.encode(FileName))
		self.checkError(None, self.lib.libprimes_setjournal(pFileName))
	


''' Class Implementation for Base
'''
class Base:
	def __init__(self, handle, wrapper):
		if not handle or not wrapper:
			raise ELibPrimesException(ErrorCodes.INVALIDPARAM)
		self._handle = handle
		self._wrapper = wrapper
	
	def __del__(self):
		self._wrapper.ReleaseInstance(self)


''' Class Implementation for Calculator
'''
class Calculator(Base):
	def __init__(self, handle, wrapper):
		Base.__init__(self, handle, wrapper)
	def GetValue(self):
		pValue = ctypes.c_uint64()
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_calculator_getvalue(self._handle, pValue))
		return pValue.value
	
	def SetValue(self, Value):
		nValue = ctypes.c_uint64(Value)
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_calculator_setvalue(self._handle, nValue))
	
	def Calculate(self):
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_calculator_calculate(self._handle))
	
	def SetProgressCallback(self, ProgressCallbackFunc):
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_calculator_setprogresscallback(self._handle, ProgressCallbackFunc))
	


''' Class Implementation for FactorizationCalculator
'''
class FactorizationCalculator(Calculator):
	def __init__(self, handle, wrapper):
		Calculator.__init__(self, handle, wrapper)
	def GetPrimeFactors(self):
		nPrimeFactorsCount = ctypes.c_uint64(0)
		nPrimeFactorsNeededCount = ctypes.c_uint64(0)
		pPrimeFactorsBuffer = (PrimeFactor*0)()
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_factorizationcalculator_getprimefactors(self._handle, nPrimeFactorsCount, nPrimeFactorsNeededCount, pPrimeFactorsBuffer))
		nPrimeFactorsCount = ctypes.c_uint64(nPrimeFactorsNeededCount.value)
		pPrimeFactorsBuffer = (PrimeFactor * nPrimeFactorsNeededCount.value)()
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_factorizationcalculator_getprimefactors(self._handle, nPrimeFactorsCount, nPrimeFactorsNeededCount, pPrimeFactorsBuffer))
		return [pPrimeFactorsBuffer[i] for i in range(nPrimeFactorsNeededCount.value)]
	


''' Class Implementation for SieveCalculator
'''
class SieveCalculator(Calculator):
	def __init__(self, handle, wrapper):
		Calculator.__init__(self, handle, wrapper)
	def GetPrimes(self):
		nPrimesCount = ctypes.c_uint64(0)
		nPrimesNeededCount = ctypes.c_uint64(0)
		pPrimesBuffer = (ctypes.c_uint64*0)()
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_sievecalculator_getprimes(self._handle, nPrimesCount, nPrimesNeededCount, pPrimesBuffer))
		nPrimesCount = ctypes.c_uint64(nPrimesNeededCount.value)
		pPrimesBuffer = (ctypes.c_uint64 * nPrimesNeededCount.value)()
		self._wrapper.checkError(self, self._wrapper.lib.libprimes_sievecalculator_getprimes(self._handle, nPrimesCount, nPrimesNeededCount, pPrimesBuffer))
		return [pPrimesBuffer[i] for i in range(nPrimesNeededCount.value)]
	
		
