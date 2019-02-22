/*++

Copyright (C) 2018 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0.

Abstract: This is an autogenerated C++ Implementation file with the basic internal
 exception type in order to allow an easy use of Prime Numbers Library

Interface version: 1.2.0

*/


#include <string>

#include "libprimes_interfaceexception.hpp"

/*************************************************************************************************************************
 Class ELibPrimesInterfaceException
**************************************************************************************************************************/
ELibPrimesInterfaceException::ELibPrimesInterfaceException(LibPrimesResult errorCode)
	: m_errorMessage("LibPrimes Error " + std::to_string (errorCode))
{
	m_errorCode = errorCode;
}

LibPrimesResult ELibPrimesInterfaceException::getErrorCode ()
{
	return m_errorCode;
}

const char * ELibPrimesInterfaceException::what () const noexcept
{
	return m_errorMessage.c_str();
}

