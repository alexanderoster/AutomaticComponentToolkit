/*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.3.1.

Abstract: This is an autogenerated C++ Header file in order to allow an easy use
 of Prime Numbers Interface

Interface version: 1.3.0

*/


#ifndef __LIBPRIMES_CPPHEADER
#define __LIBPRIMES_CPPHEADER

#include "libprimes.h"
#include <string>
#include <memory>
#include <vector>
#include <exception>

namespace LibPrimes {

/*************************************************************************************************************************
 Forward Declaration of all classes 
**************************************************************************************************************************/

class CLibPrimesBaseClass;
class CLibPrimesCalculator;
class CLibPrimesFactorizationCalculator;
class CLibPrimesSieveCalculator;

/*************************************************************************************************************************
 Declaration of shared pointer types 
**************************************************************************************************************************/

typedef std::shared_ptr<CLibPrimesBaseClass> PLibPrimesBaseClass;
typedef std::shared_ptr<CLibPrimesCalculator> PLibPrimesCalculator;
typedef std::shared_ptr<CLibPrimesFactorizationCalculator> PLibPrimesFactorizationCalculator;
typedef std::shared_ptr<CLibPrimesSieveCalculator> PLibPrimesSieveCalculator;
		 
/*************************************************************************************************************************
 Class ELibPrimesException 
**************************************************************************************************************************/
class ELibPrimesException : public std::exception {
	protected:
		/**
		* Error code for the Exception.
		*/
		LibPrimesResult m_errorCode;
		/**
		* Error message for the Exception.
		*/
		std::string m_errorMessage;

	public:
		/**
		* Exception Constructor.
		*/
		ELibPrimesException (LibPrimesResult errorCode);

		/**
		* Returns error code
		*/
		LibPrimesResult getErrorCode ();

		/**
		* Returns error message
		*/
		const char* what () const noexcept;

};

/*************************************************************************************************************************
 Class CLibPrimesInputVector
**************************************************************************************************************************/
template <typename T>
class CLibPrimesInputVector {
private:
	
	const T* m_data;
	size_t m_size;
	
public:
	
	CLibPrimesInputVector( const std::vector<T>& vec)
		: m_data( vec.data() ), m_size( vec.size() )
	{
	}
	
	CLibPrimesInputVector( const T* in_data, size_t in_size)
		: m_data( in_data ), m_size(in_size )
	{
	}
	
	const T* data() const
	{
		return m_data;
	}
	
	size_t size() const
	{
		return m_size;
	}
	
};


/*************************************************************************************************************************
 Class CLibPrimesBaseClass 
**************************************************************************************************************************/
class CLibPrimesBaseClass {
protected:
	/* Handle to Instance in library*/
	LibPrimesHandle m_pHandle;

	/* Checks for an Error code and raises Exceptions */
	void CheckError(LibPrimesResult nResult);
public:

	/**
	* CLibPrimesBaseClass::CLibPrimesBaseClass - Constructor for Base class.
	*/
	CLibPrimesBaseClass(LibPrimesHandle pHandle);

	/**
	* CLibPrimesBaseClass::~CLibPrimesBaseClass - Destructor for Base class.
	*/
	virtual ~CLibPrimesBaseClass();

	/**
	* CLibPrimesBaseClass::GetHandle - Returns handle to instance.
	*/
	LibPrimesHandle GetHandle();
};
		 
/*************************************************************************************************************************
 Class CLibPrimesCalculator 
**************************************************************************************************************************/
class CLibPrimesCalculator : public CLibPrimesBaseClass {
public:
	
	/**
	* CLibPrimesCalculator::CLibPrimesCalculator - Constructor for Calculator class.
	*/
	CLibPrimesCalculator (LibPrimesHandle pHandle);

	/**
	* CLibPrimesCalculator::GetValue - Returns the current value of this Calculator
	* @return The current value of this Calculator
	*/
	LibPrimes_uint64 GetValue ();

	/**
	* CLibPrimesCalculator::SetValue - Sets the value to be factorized
	* @param[in] nValue - The value to be factorized
	*/
	void SetValue (const LibPrimes_uint64 nValue);

	/**
	* CLibPrimesCalculator::SetProgressCallback - Sets the progress callback function
	* @param[in] pProgressCallback - The progress callback
	*/
	void SetProgressCallback (const LibPrimesProgressCallback pProgressCallback);

	/**
	* CLibPrimesCalculator::Calculate - Performs the specific calculation of this Calculator
	*/
	void Calculate ();
};
		 
/*************************************************************************************************************************
 Class CLibPrimesFactorizationCalculator 
**************************************************************************************************************************/
class CLibPrimesFactorizationCalculator : public CLibPrimesCalculator {
public:
	
	/**
	* CLibPrimesFactorizationCalculator::CLibPrimesFactorizationCalculator - Constructor for FactorizationCalculator class.
	*/
	CLibPrimesFactorizationCalculator (LibPrimesHandle pHandle);

	/**
	* CLibPrimesFactorizationCalculator::GetPrimeFactors - Returns the prime factors of this number (with multiplicity)
	* @param[out] PrimeFactorsBuffer - The prime factors of this number
	*/
	void GetPrimeFactors (std::vector<sLibPrimesPrimeFactor> & PrimeFactorsBuffer);

	/**
	* CLibPrimesFactorizationCalculator::CheckPrimeFactors - Checks, whether a list of prime factors (with multiplicity) is the prime factor decomposistion of the calculator's value
	* @param[in] PrimeFactorsBuffer - 
	* @return Do the prime factors decompose this calculator's value
	*/
	bool CheckPrimeFactors (const CLibPrimesInputVector<sLibPrimesPrimeFactor> & PrimeFactorsBuffer);
};
		 
/*************************************************************************************************************************
 Class CLibPrimesSieveCalculator 
**************************************************************************************************************************/
class CLibPrimesSieveCalculator : public CLibPrimesCalculator {
public:
	
	/**
	* CLibPrimesSieveCalculator::CLibPrimesSieveCalculator - Constructor for SieveCalculator class.
	*/
	CLibPrimesSieveCalculator (LibPrimesHandle pHandle);

	/**
	* CLibPrimesSieveCalculator::GetPrimes - Returns all prime numbers lower or equal to the sieve's value
	* @param[out] PrimesBuffer - The primes lower or equal to the sieve's value
	*/
	void GetPrimes (std::vector<LibPrimes_uint64> & PrimesBuffer);
};
		 
/*************************************************************************************************************************
 Class CLibPrimesWrapper 
**************************************************************************************************************************/
class CLibPrimesWrapper {
public:
	static void CheckError(LibPrimesHandle handle, LibPrimesResult nResult);

	/**
	* CLibPrimesWrapper::CreateFactorizationCalculator - Creates a new FactorizationCalculator instance
	* @return New FactorizationCalculator instance
	*/
	static PLibPrimesFactorizationCalculator CreateFactorizationCalculator ();

	/**
	* CLibPrimesWrapper::CreateSieveCalculator - Creates a new SieveCalculator instance
	* @return New SieveCalculator instance
	*/
	static PLibPrimesSieveCalculator CreateSieveCalculator ();

	/**
	* CLibPrimesWrapper::ReleaseInstance - Releases the memory of an Instance
	* @param[in] pInstance - Instance Handle
	*/
	static void ReleaseInstance (CLibPrimesBaseClass * pInstance);

	/**
	* CLibPrimesWrapper::GetLibraryVersion - retrieves the current version of the library.
	* @param[out] nMajor - returns the major version of the library
	* @param[out] nMinor - returns the minor version of the library
	* @param[out] nMicro - returns the micro version of the library
	*/
	static void GetLibraryVersion (LibPrimes_uint32 & nMajor, LibPrimes_uint32 & nMinor, LibPrimes_uint32 & nMicro);

	/**
	* CLibPrimesWrapper::SetJournal - Handles Library Journaling
	* @param[in] sFileName - Journal FileName
	*/
	static void SetJournal (const std::string & sFileName);
};

};

#endif // __LIBPRIMES_CPPHEADER

