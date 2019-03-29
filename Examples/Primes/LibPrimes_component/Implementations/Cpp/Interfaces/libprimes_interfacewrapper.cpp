/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0-develop5.

Abstract: This is an autogenerated C++ implementation file in order to allow easy
development of Prime Numbers Library. The functions in this file need to be implemented. It needs to be generated only once.

Interface version: 1.2.0

*/

#include "libprimes_abi.hpp"
#include "libprimes_interfaces.hpp"
#include "libprimes_interfaceexception.hpp"
#include "libprimes_interfacejournal.hpp"

using namespace LibPrimes::Impl;

PLibPrimesInterfaceJournal m_GlobalJournal;

LibPrimesResult handleLibPrimesException(IBase * pIBaseClass, ELibPrimesInterfaceException & Exception, CLibPrimesInterfaceJournalEntry * pJournalEntry = nullptr)
{
	LibPrimesResult errorCode = Exception.getErrorCode();

	if (pJournalEntry != nullptr)
		pJournalEntry->writeError(errorCode);

	if (pIBaseClass != nullptr)
		pIBaseClass->RegisterErrorMessage(Exception.what());

	return errorCode;
}

LibPrimesResult handleStdException(IBase * pIBaseClass, std::exception & Exception, CLibPrimesInterfaceJournalEntry * pJournalEntry = nullptr)
{
	LibPrimesResult errorCode = LIBPRIMES_ERROR_GENERICEXCEPTION;

	if (pJournalEntry != nullptr)
		pJournalEntry->writeError(errorCode);

	if (pIBaseClass != nullptr)
		pIBaseClass->RegisterErrorMessage(Exception.what());

	return errorCode;
}

LibPrimesResult handleUnhandledException(IBase * pIBaseClass, CLibPrimesInterfaceJournalEntry * pJournalEntry = nullptr)
{
	LibPrimesResult errorCode = LIBPRIMES_ERROR_GENERICEXCEPTION;

	if (pJournalEntry != nullptr)
		pJournalEntry->writeError(errorCode);

	if (pIBaseClass != nullptr)
		pIBaseClass->RegisterErrorMessage("Unhandled Exception");

	return errorCode;
}



/*************************************************************************************************************************
 Class implementation for Base
**************************************************************************************************************************/

/*************************************************************************************************************************
 Class implementation for Calculator
**************************************************************************************************************************/
LibPrimesResult libprimes_calculator_getvalue(LibPrimes_Calculator pCalculator, LibPrimes_uint64 * pValue)
{
	IBase* pIBaseClass = (IBase *)pCalculator;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginClassMethod(pCalculator, "Calculator", "GetValue");
		}

		if (pValue == nullptr)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		ICalculator* pICalculator = dynamic_cast<ICalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);


		*pValue = pICalculator->GetValue();


		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->addUInt64Result("Value", *pValue);
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_calculator_getself(LibPrimes_Calculator pCalculator, LibPrimes_Calculator * pValue)
{
	IBase* pIBaseClass = (IBase *)pCalculator;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginClassMethod(pCalculator, "Calculator", "GetSelf");
		}

		if (pValue == nullptr)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		ICalculator* pICalculator = dynamic_cast<ICalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);

		IBase* pBaseValue(nullptr);

		pBaseValue = pICalculator->GetSelf();

		*pValue = (IBase*)(pBaseValue);

		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->addHandleResult("Value", *pValue);
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_calculator_setvalue(LibPrimes_Calculator pCalculator, LibPrimes_uint64 nValue)
{
	IBase* pIBaseClass = (IBase *)pCalculator;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginClassMethod(pCalculator, "Calculator", "SetValue");
			pJournalEntry->addUInt64Parameter("Value", nValue);
		}


		ICalculator* pICalculator = dynamic_cast<ICalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);


		pICalculator->SetValue(nValue);


		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_calculator_calculate(LibPrimes_Calculator pCalculator)
{
	IBase* pIBaseClass = (IBase *)pCalculator;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginClassMethod(pCalculator, "Calculator", "Calculate");
		}


		ICalculator* pICalculator = dynamic_cast<ICalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);


		pICalculator->Calculate();


		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_calculator_setprogresscallback(LibPrimes_Calculator pCalculator, LibPrimesProgressCallback pProgressCallback)
{
	IBase* pIBaseClass = (IBase *)pCalculator;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginClassMethod(pCalculator, "Calculator", "SetProgressCallback");
		}


		ICalculator* pICalculator = dynamic_cast<ICalculator*>(pIBaseClass);
		if (!pICalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);


		pICalculator->SetProgressCallback(pProgressCallback);


		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}


/*************************************************************************************************************************
 Class implementation for FactorizationCalculator
**************************************************************************************************************************/
LibPrimesResult libprimes_factorizationcalculator_getprimefactors(LibPrimes_FactorizationCalculator pFactorizationCalculator, const LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64* pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer)
{
	IBase* pIBaseClass = (IBase *)pFactorizationCalculator;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginClassMethod(pFactorizationCalculator, "FactorizationCalculator", "GetPrimeFactors");
		}

		if ((!pPrimeFactorsBuffer) && !(pPrimeFactorsNeededCount))
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		IFactorizationCalculator* pIFactorizationCalculator = dynamic_cast<IFactorizationCalculator*>(pIBaseClass);
		if (!pIFactorizationCalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);


		pIFactorizationCalculator->GetPrimeFactors(nPrimeFactorsBufferSize, pPrimeFactorsNeededCount, pPrimeFactorsBuffer);


		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}


/*************************************************************************************************************************
 Class implementation for SieveCalculator
**************************************************************************************************************************/
LibPrimesResult libprimes_sievecalculator_getprimes(LibPrimes_SieveCalculator pSieveCalculator, const LibPrimes_uint64 nPrimesBufferSize, LibPrimes_uint64* pPrimesNeededCount, LibPrimes_uint64 * pPrimesBuffer)
{
	IBase* pIBaseClass = (IBase *)pSieveCalculator;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginClassMethod(pSieveCalculator, "SieveCalculator", "GetPrimes");
		}

		if ((!pPrimesBuffer) && !(pPrimesNeededCount))
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		ISieveCalculator* pISieveCalculator = dynamic_cast<ISieveCalculator*>(pIBaseClass);
		if (!pISieveCalculator)
			throw ELibPrimesInterfaceException(LIBPRIMES_ERROR_INVALIDCAST);


		pISieveCalculator->GetPrimes(nPrimesBufferSize, pPrimesNeededCount, pPrimesBuffer);


		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}


/*************************************************************************************************************************
 Global functions implementation
**************************************************************************************************************************/
LibPrimesResult libprimes_getlasterror(LibPrimes_Base pInstance, const LibPrimes_uint32 nErrorMessageBufferSize, LibPrimes_uint32* pErrorMessageNeededChars, char * pErrorMessageBuffer, bool * pHasError)
{
	IBase* pIBaseClass = nullptr;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginStaticFunction("GetLastError");
			pJournalEntry->addHandleParameter("Instance", pInstance);
		}

		if ( (!pErrorMessageBuffer) && !(pErrorMessageNeededChars) )
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);
		if (pHasError == nullptr)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		IBase* pIBaseClassInstance = (IBase *)pInstance;
		IBase* pIInstance = dynamic_cast<IBase*>(pIBaseClassInstance);
		if (!pIInstance)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDCAST);

		std::string sErrorMessage("");

		*pHasError = CWrapper::GetLastError(pIInstance, sErrorMessage);

		if (pErrorMessageNeededChars) 
			*pErrorMessageNeededChars = (LibPrimes_uint32) sErrorMessage.size();
		if (pErrorMessageBuffer) {
			if (sErrorMessage.size() >= nErrorMessageBufferSize)
				throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_BUFFERTOOSMALL);
			for (int iErrorMessage = 0; iErrorMessage < sErrorMessage.size(); iErrorMessage++)
				pErrorMessageBuffer[iErrorMessage] = sErrorMessage[iErrorMessage];
		}

		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->addStringResult("ErrorMessage", sErrorMessage.c_str());
			pJournalEntry->addBooleanResult("HasError", *pHasError);
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_releaseinstance(LibPrimes_Base pInstance)
{
	IBase* pIBaseClass = nullptr;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginStaticFunction("ReleaseInstance");
			pJournalEntry->addHandleParameter("Instance", pInstance);
		}


		IBase* pIBaseClassInstance = (IBase *)pInstance;
		IBase* pIInstance = dynamic_cast<IBase*>(pIBaseClassInstance);
		if (!pIInstance)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDCAST);


		CWrapper::ReleaseInstance(pIInstance);


		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_getlibraryversion(LibPrimes_uint32 * pMajor, LibPrimes_uint32 * pMinor, LibPrimes_uint32 * pMicro, const LibPrimes_uint32 nPreReleaseInfoBufferSize, LibPrimes_uint32* pPreReleaseInfoNeededChars, char * pPreReleaseInfoBuffer, const LibPrimes_uint32 nBuildInfoBufferSize, LibPrimes_uint32* pBuildInfoNeededChars, char * pBuildInfoBuffer)
{
	IBase* pIBaseClass = nullptr;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginStaticFunction("GetLibraryVersion");
		}

		if (!pMajor)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);
		if (!pMinor)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);
		if (!pMicro)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);
		if ( (!pPreReleaseInfoBuffer) && !(pPreReleaseInfoNeededChars) )
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);
		if ( (!pBuildInfoBuffer) && !(pBuildInfoNeededChars) )
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		std::string sPreReleaseInfo("");
		std::string sBuildInfo("");

		CWrapper::GetLibraryVersion(*pMajor, *pMinor, *pMicro, sPreReleaseInfo, sBuildInfo);

		if (pPreReleaseInfoNeededChars) 
			*pPreReleaseInfoNeededChars = (LibPrimes_uint32) sPreReleaseInfo.size();
		if (pPreReleaseInfoBuffer) {
			if (sPreReleaseInfo.size() >= nPreReleaseInfoBufferSize)
				throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_BUFFERTOOSMALL);
			for (int iPreReleaseInfo = 0; iPreReleaseInfo < sPreReleaseInfo.size(); iPreReleaseInfo++)
				pPreReleaseInfoBuffer[iPreReleaseInfo] = sPreReleaseInfo[iPreReleaseInfo];
		}
		if (pBuildInfoNeededChars) 
			*pBuildInfoNeededChars = (LibPrimes_uint32) sBuildInfo.size();
		if (pBuildInfoBuffer) {
			if (sBuildInfo.size() >= nBuildInfoBufferSize)
				throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_BUFFERTOOSMALL);
			for (int iBuildInfo = 0; iBuildInfo < sBuildInfo.size(); iBuildInfo++)
				pBuildInfoBuffer[iBuildInfo] = sBuildInfo[iBuildInfo];
		}

		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->addUInt32Result("Major", *pMajor);
			pJournalEntry->addUInt32Result("Minor", *pMinor);
			pJournalEntry->addUInt32Result("Micro", *pMicro);
			pJournalEntry->addStringResult("PreReleaseInfo", sPreReleaseInfo.c_str());
			pJournalEntry->addStringResult("BuildInfo", sBuildInfo.c_str());
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_createfactorizationcalculator(LibPrimes_FactorizationCalculator * pInstance)
{
	IBase* pIBaseClass = nullptr;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginStaticFunction("CreateFactorizationCalculator");
		}

		if (pInstance == nullptr)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		IBase* pBaseInstance(nullptr);

		pBaseInstance = CWrapper::CreateFactorizationCalculator();

		*pInstance = (IBase*)(pBaseInstance);

		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->addHandleResult("Instance", *pInstance);
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_createsievecalculator(LibPrimes_SieveCalculator * pInstance)
{
	IBase* pIBaseClass = nullptr;

	PLibPrimesInterfaceJournalEntry pJournalEntry;
	try {
		if (m_GlobalJournal.get() != nullptr)  {
			pJournalEntry = m_GlobalJournal->beginStaticFunction("CreateSieveCalculator");
		}

		if (pInstance == nullptr)
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		IBase* pBaseInstance(nullptr);

		pBaseInstance = CWrapper::CreateSieveCalculator();

		*pInstance = (IBase*)(pBaseInstance);

		if (pJournalEntry.get() != nullptr) {
			pJournalEntry->addHandleResult("Instance", *pInstance);
			pJournalEntry->writeSuccess();
		}

		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception, pJournalEntry.get());
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException, pJournalEntry.get());
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass, pJournalEntry.get());
	}
}

LibPrimesResult libprimes_setjournal(const char * pFileName)
{
	IBase* pIBaseClass = nullptr;

	try {
		if (pFileName == nullptr) 
			throw ELibPrimesInterfaceException (LIBPRIMES_ERROR_INVALIDPARAM);

		std::string sFileName(pFileName);

		m_GlobalJournal = nullptr;
		if (sFileName != "") {
			m_GlobalJournal = std::make_shared<CLibPrimesInterfaceJournal> (sFileName);
		}


		return LIBPRIMES_SUCCESS;
	}
	catch (ELibPrimesInterfaceException & Exception) {
		return handleLibPrimesException(pIBaseClass, Exception);
	}
	catch (std::exception & StdException) {
		return handleStdException(pIBaseClass, StdException);
	}
	catch (...) {
		return handleUnhandledException(pIBaseClass);
	}
}


