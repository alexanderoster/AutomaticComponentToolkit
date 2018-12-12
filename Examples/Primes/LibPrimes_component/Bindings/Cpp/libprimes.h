/*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.3.1.

Abstract: This is an autogenerated plain C Header file in order to allow an easy
 use of Prime Numbers Interface

Interface version: 1.3.0

*/

#ifndef __LIBPRIMES_HEADER
#define __LIBPRIMES_HEADER

#ifdef __LIBPRIMES_EXPORTS
#ifdef WIN32
#define LIBPRIMES_DECLSPEC __declspec (dllexport)
#else // WIN32
#define LIBPRIMES_DECLSPEC __attribute__((visibility("default")))
#endif // WIN32
#else // __LIBPRIMES_EXPORTS
#define LIBPRIMES_DECLSPEC
#endif // __LIBPRIMES_EXPORTS

#include "libprimes_types.h"

extern "C" {

/*************************************************************************************************************************
 Class definition for Calculator
**************************************************************************************************************************/

/**
* Returns the current value of this Calculator
*
* @param[in] pCalculator - Calculator instance.
* @param[out] pValue - The current value of this Calculator
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_calculator_getvalue (LibPrimes_Calculator pCalculator, LibPrimes_uint64 * pValue);

/**
* Sets the value to be factorized
*
* @param[in] pCalculator - Calculator instance.
* @param[in] nValue - The value to be factorized
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_calculator_setvalue (LibPrimes_Calculator pCalculator, LibPrimes_uint64 nValue);

/**
* Sets the progress callback function
*
* @param[in] pCalculator - Calculator instance.
* @param[in] pProgressCallback - The progress callback
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_calculator_setprogresscallback (LibPrimes_Calculator pCalculator, LibPrimesProgressCallback pProgressCallback);

/**
* Performs the specific calculation of this Calculator
*
* @param[in] pCalculator - Calculator instance.
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_calculator_calculate (LibPrimes_Calculator pCalculator);

/*************************************************************************************************************************
 Class definition for FactorizationCalculator
**************************************************************************************************************************/

/**
* Returns the prime factors of this number (with multiplicity)
*
* @param[in] pFactorizationCalculator - FactorizationCalculator instance.
* @param[in] nPrimeFactorsBufferSize - Number of elements in buffer
* @param[out] pPrimeFactorsNeededCount - will be filled with the count of the written elements, or needed buffer size.
* @param[out] pPrimeFactorsBuffer - PrimeFactor buffer of The prime factors of this number
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_factorizationcalculator_getprimefactors (LibPrimes_FactorizationCalculator pFactorizationCalculator, const LibPrimes_uint64 nPrimeFactorsBufferSize, LibPrimes_uint64* pPrimeFactorsNeededCount, sLibPrimesPrimeFactor * pPrimeFactorsBuffer);

/**
* Checks, whether a list of prime factors (with multiplicity) is the prime factor decomposistion of the calculator's value
*
* @param[in] pFactorizationCalculator - FactorizationCalculator instance.
* @param[in] nPrimeFactorsBufferSize - Number of elements in buffer
* @param[in] pPrimeFactorsBuffer - PrimeFactor buffer of 
* @param[out] pAreEqual - Do the prime factors decompose this calculator's value
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_factorizationcalculator_checkprimefactors (LibPrimes_FactorizationCalculator pFactorizationCalculator, LibPrimes_uint64 nPrimeFactorsBufferSize, const sLibPrimesPrimeFactor * pPrimeFactorsBuffer, bool * pAreEqual);

/*************************************************************************************************************************
 Class definition for SieveCalculator
**************************************************************************************************************************/

/**
* Returns all prime numbers lower or equal to the sieve's value
*
* @param[in] pSieveCalculator - SieveCalculator instance.
* @param[in] nPrimesBufferSize - Number of elements in buffer
* @param[out] pPrimesNeededCount - will be filled with the count of the written elements, or needed buffer size.
* @param[out] pPrimesBuffer - uint64 buffer of The primes lower or equal to the sieve's value
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_sievecalculator_getprimes (LibPrimes_SieveCalculator pSieveCalculator, const LibPrimes_uint64 nPrimesBufferSize, LibPrimes_uint64* pPrimesNeededCount, LibPrimes_uint64 * pPrimesBuffer);

/*************************************************************************************************************************
 Global functions
**************************************************************************************************************************/

/**
* Creates a new FactorizationCalculator instance
*
* @param[out] pInstance - New FactorizationCalculator instance
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_createfactorizationcalculator (LibPrimes_FactorizationCalculator * pInstance);

/**
* Creates a new SieveCalculator instance
*
* @param[out] pInstance - New SieveCalculator instance
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_createsievecalculator (LibPrimes_SieveCalculator * pInstance);

/**
* Releases the memory of an Instance
*
* @param[in] pInstance - Instance Handle
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_releaseinstance (LibPrimes_BaseClass pInstance);

/**
* retrieves the current version of the library.
*
* @param[out] pMajor - returns the major version of the library
* @param[out] pMinor - returns the minor version of the library
* @param[out] pMicro - returns the micro version of the library
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_getlibraryversion (LibPrimes_uint32 * pMajor, LibPrimes_uint32 * pMinor, LibPrimes_uint32 * pMicro);

/**
* Handles Library Journaling
*
* @param[in] pFileName - Journal FileName
* @return error code or 0 (success)
*/
LIBPRIMES_DECLSPEC LibPrimesResult libprimes_setjournal (const char * pFileName);

}

#endif // __LIBPRIMES_HEADER

