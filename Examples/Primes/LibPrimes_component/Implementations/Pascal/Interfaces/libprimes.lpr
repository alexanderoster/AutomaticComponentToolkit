(*++

Copyright (C) 2018 Automatic Component Toolkit Developers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.3.1.

Abstract: This is an autogenerated Pascal project file in order to allow easy
development of Prime Numbers Interface.

Interface version: 1.3.0

*)

{$MODE DELPHI}
library libprimes;

uses
{$IFDEF UNIX}
	cthreads,
{$ENDIF UNIX}
	syncobjs,
	libprimes_types,
	libprimes_exports,
	Classes,
	sysutils;

exports
	libprimes_calculator_getvalue,
	libprimes_calculator_setvalue,
	libprimes_calculator_setprogresscallback,
	libprimes_calculator_calculate,
	libprimes_factorizationcalculator_getprimefactors,
	libprimes_factorizationcalculator_checkprimefactors,
	libprimes_sievecalculator_getprimes,
	libprimes_createfactorizationcalculator,
	libprimes_createsievecalculator,
	libprimes_releaseinstance,
	libprimes_getlibraryversion,
	libprimes_setjournal;

begin

end.
