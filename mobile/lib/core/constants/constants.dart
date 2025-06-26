// lib/app_constants.dart
import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

// Colors
const Color kBackgroundColor = Color(0xFFF4F0ED);
const Color kPrimaryButtonColor = Color(0xFF946B4A);
const Color kSecondaryButtonColor = Color(0xFFFFFFFF);
//const Color kLandingPageColor = Color(0xFFD9EAE3); // Approximation of the textured green
const Color kTextColor = Color(0xFF876143);
const Color kLinkColor = Color(0XFF2A55C4);

// Text Styles
final TextStyle kTitleStyle = GoogleFonts.poppins(
  fontSize: 28,
  fontWeight: FontWeight.bold,
  color: kTextColor,
);

final TextStyle kSubtitleStyle = GoogleFonts.poppins(
  fontSize: 16,
  color: Color(0xFF989898),
);

final TextStyle kFormLabelStyle = GoogleFonts.poppins(
  fontSize: 14,
  color: Color(0xFF876143),
);

final TextStyle kBodyTextStyle = GoogleFonts.poppins(
  fontSize: 16,
  color: kTextColor,
);

final TextStyle kButtonTextStyle = GoogleFonts.poppins(
  fontSize: 16,
  fontWeight: FontWeight.w600,
  color: Colors.white,
);