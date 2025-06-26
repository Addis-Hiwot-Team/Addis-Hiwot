// lib/verify_email_page.dart
import 'package:flutter/material.dart';
import 'package:mobile/core/constants/constants.dart';
import 'package:pinput/pinput.dart';

class VerifyEmailPage extends StatelessWidget {
  const VerifyEmailPage({super.key});

  @override
  Widget build(BuildContext context) {
    final defaultPinTheme = PinTheme(
      width: 56,
      height: 56,
      textStyle: const TextStyle(fontSize: 20, color: kTextColor, fontWeight: FontWeight.w600),
      decoration: BoxDecoration(
        border: Border.all(color: Colors.grey.shade400),
        borderRadius: BorderRadius.circular(12),
      ),
    );

    return Scaffold(
      backgroundColor: kBackgroundColor,
      appBar: AppBar(
        backgroundColor: kBackgroundColor,
        elevation: 0,
        leading: IconButton(
          icon: const Icon(Icons.arrow_back, color: kTextColor),
          onPressed: () => Navigator.of(context).pop(),
        ),
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.symmetric(horizontal: 24.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.stretch,
            children: [
              const SizedBox(height: 20),
              Text('Check Your Email', style: kTitleStyle),
              const SizedBox(height: 16),
              RichText(
                text: TextSpan(
                  style: kSubtitleStyle,
                  children: const [
                    TextSpan(text: 'We sent a reset link to '),
                    TextSpan(
                      text: 'contact@decode...com ',
                      style: TextStyle(color: kTextColor, fontWeight: FontWeight.bold),
                    ),
                    TextSpan(text: 'enter 5 digit code that mentioned in the email'),
                  ],
                ),
              ),
              const SizedBox(height: 40),
              Pinput(
                length: 5,
                defaultPinTheme: defaultPinTheme,
                focusedPinTheme: defaultPinTheme.copyWith(
                  decoration: defaultPinTheme.decoration!.copyWith(
                    border: Border.all(color: kPrimaryButtonColor),
                  ),
                ),
                submittedPinTheme: defaultPinTheme,
                validator: (s) {
                  return s == '86300' ? null : 'Pin is incorrect';
                },
                pinputAutovalidateMode: PinputAutovalidateMode.onSubmit,
                showCursor: true,
                onCompleted: (pin) => print(pin),
              ),
              const SizedBox(height: 40),
              ElevatedButton(
                style: ElevatedButton.styleFrom(
                  backgroundColor: kPrimaryButtonColor,
                  padding: const EdgeInsets.symmetric(vertical: 16),
                  shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
                ),
                onPressed: () {
                  Navigator.pushNamed(context, '/reset_password');
                },
                child: Text('Verify Code', style: kButtonTextStyle),
              ),
              const SizedBox(height: 20),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  const Text("Haven't got the email yet?"),
                  TextButton(
                    onPressed: () {},
                    child: const Text('Resend email',
                        style: TextStyle(
                            color: kLinkColor, fontWeight: FontWeight.bold)),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
}