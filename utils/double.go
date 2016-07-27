package utils

/*
   if (d1 < d2)
       return -1;           // Neither val is NaN, thisVal is smaller
   if (d1 > d2)
       return 1;            // Neither val is NaN, thisVal is larger

   // Cannot use doubleToRawLongBits because of possibility of NaNs.
   long thisBits    = Double.doubleToLongBits(d1);
   long anotherBits = Double.doubleToLongBits(d2);

   return (thisBits == anotherBits ?  0 : // Values are equal
           (thisBits < anotherBits ? -1 : // (-0.0, 0.0) or (!NaN, NaN)
            1));
*/

func CompareTo(d1 float64, d2 float64) int {
	panic("not implement")
}

/**
 * Returns a representation of the specified floating-point value
 * according to the IEEE 754 floating-point "double
 * format" bit layout.
 *
 * <p>Bit 63 (the bit that is selected by the mask
 * {@code 0x8000000000000000L}) represents the sign of the
 * floating-point number. Bits
 * 62-52 (the bits that are selected by the mask
 * {@code 0x7ff0000000000000L}) represent the exponent. Bits 51-0
 * (the bits that are selected by the mask
 * {@code 0x000fffffffffffffL}) represent the significand
 * (sometimes called the mantissa) of the floating-point number.
 *
 * <p>If the argument is positive infinity, the result is
 * {@code 0x7ff0000000000000L}.
 *
 * <p>If the argument is negative infinity, the result is
 * {@code 0xfff0000000000000L}.
 *
 * <p>If the argument is NaN, the result is
 * {@code 0x7ff8000000000000L}.
 *
 * <p>In all cases, the result is a {@code long} integer that, when
 * given to the {@link #longBitsToDouble(long)} method, will produce a
 * floating-point value the same as the argument to
 * {@code doubleToLongBits} (except all NaN values are
 * collapsed to a single "canonical" NaN value).
 *
 * @param   value   a {@code double} precision floating-point number.
 * @return the bits that represent the floating-point number.
 */
/*
   public static long doubleToLongBits(double value) {
       long result = doubleToRawLongBits(value);
       // Check for NaN based on values of bit fields, maximum
       // exponent and nonzero significand.
       if ( ((result & DoubleConsts.EXP_BIT_MASK) ==
             DoubleConsts.EXP_BIT_MASK) &&
            (result & DoubleConsts.SIGNIF_BIT_MASK) != 0L)
           result = 0x7ff8000000000000L;
       return result;
   }*/

func Float64ToLongBits(value float64) int64 {
	panic("not implement")
}
