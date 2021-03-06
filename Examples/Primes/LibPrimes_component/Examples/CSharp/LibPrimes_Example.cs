/*++

Copyright (C) 2019 PrimeDevelopers

All rights reserved.

This file has been generated by the Automatic Component Toolkit (ACT) version 1.5.0.

Abstract: This is an autogenerated CSharp application that demonstrates the
 usage of the CSharp bindings of Prime Numbers Library

Interface version: 1.2.0

*/


using System;
namespace LibPrimes_Example
{
	class LibPrimes_Example
	{
		static void Main()
		{
			UInt32 nMajor, nMinor, nMicro;
			LibPrimes.Wrapper.GetVersion(out nMajor, out nMinor, out nMicro);
			string versionString = string.Format("LibPrimes.version = {0}.{1}.{2}", nMajor, nMinor, nMicro);
			Console.WriteLine(versionString);

            LibPrimes.CFactorizationCalculator factorization = LibPrimes.Wrapper.CreateFactorizationCalculator();
            factorization.SetValue(735);
            factorization.Calculate();
            LibPrimes.sPrimeFactor[] aPrimeFactors;
            factorization.GetPrimeFactors(out aPrimeFactors);

            Console.Write(string.Format("{0} = 1 ", factorization.GetValue()));
            foreach (LibPrimes.sPrimeFactor pF in aPrimeFactors)
            {
                Console.Write(string.Format("* {0}^{1} ", pF.Prime, pF.Multiplicity));
            }
            Console.WriteLine();

            Console.WriteLine("Press any key to exit.");
			Console.ReadKey();
		}
	}
}

