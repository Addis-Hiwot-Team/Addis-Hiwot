import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'app_colors.dart';

class AppTheme {
  static final ThemeData lightTheme = ThemeData(
    brightness: Brightness.light,
    primaryColor: AppColors.primary,
    scaffoldBackgroundColor: AppColors.background,
    appBarTheme: AppBarTheme(
      backgroundColor: AppColors.primary,
      foregroundColor: Colors.white,
    ),
    textTheme: GoogleFonts.epilogueTextTheme(),
    colorScheme: ColorScheme.light(
      primary: AppColors.primary,
      secondary: AppColors.secondary,
      error: AppColors.error,
    ),
  );

  static final ThemeData darkTheme = ThemeData(
    brightness: Brightness.dark,
    primaryColor: AppColors.primary,
    scaffoldBackgroundColor: Colors.black,
    appBarTheme: AppBarTheme(
      backgroundColor: Colors.black,
      foregroundColor: Colors.white,
    ),
    textTheme: GoogleFonts.epilogueTextTheme(ThemeData.dark().textTheme),
    colorScheme: ColorScheme.dark(
      primary: AppColors.primary,
      secondary: AppColors.secondary,
      error: AppColors.error,
    ),
  );
}
