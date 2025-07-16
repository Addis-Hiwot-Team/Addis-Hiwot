// lib/screens/daily_checkin_screen.dart

import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:mobile/core/constants/constants.dart';
import 'package:mobile/features/daily_goals/presentation/pages/goal_progress.dart';
import 'package:provider/provider.dart';
import '../providers/daily_checkin_provider.dart';

class DailyCheckInScreen extends StatelessWidget {
  DailyCheckInScreen({super.key});

  final List<String> moods = ['😊', '🙂', '😡', '😢', '😭'];

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => DailyCheckInProvider(),
      child: Consumer<DailyCheckInProvider>(
        builder: (context, provider, _) {
          return Scaffold(
            backgroundColor: kBackgroundColor,
            appBar: AppBar(
              backgroundColor: kBackgroundColor,
              leading: IconButton(
                icon: const Icon(Icons.arrow_back_ios_new),
                onPressed: () => {},
              ),
            ),
            body: ListView(
              padding: const EdgeInsets.all(20.0),
              children: [
                if (!provider.isSaved) ...[
                  Text(
                    'Good Afternoon,\nKidist!',
                    style: kTitleStyle,
                  ),
                  const SizedBox(height: 30),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Text(
                        'How are you feeling today?',
                        style: kBodyTextStyle.copyWith(fontWeight: FontWeight.w600),
                      ),
                    ],
                  ),
                  const SizedBox(height: 20),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceAround,
                    children: List.generate(moods.length, (index) {
                      return GestureDetector(
                        onTap: () {
                          provider.selectMood(index);
                        },
                        child: Opacity(
                          opacity: provider.selectedMoodIndex == null || provider.selectedMoodIndex == index ? 1.0 : 0.5,
                          child: Container(
                            padding: const EdgeInsets.all(8),
                            decoration: BoxDecoration(
                              color: provider.selectedMoodIndex == index ? kSearchBarBg : Colors.transparent,
                              borderRadius: BorderRadius.circular(12),
                            ),
                            child: Text(moods[index], style: const TextStyle(fontSize: 32)),
                          ),
                        ),
                      );
                    }),
                  ),
                  const SizedBox(height: 30),
                  Text(
                    'Write down what you feel.',
                    style: kBodyTextStyle.copyWith(fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(height: 10),
                  Container(
                    decoration: BoxDecoration(
                      color: Color(0xFFE5E5EA), // light beige
                      borderRadius: BorderRadius.circular(14),
                      border: Border.all(color: kPrimaryBrown), // subtle brown border
                    ),
                    child: TextField(
                      maxLines: 8,
                      style: kBodyTextStyle.copyWith(color: kPrimaryBrown),
                      onChanged: provider.setNote,
                      decoration: InputDecoration(
                        contentPadding: const EdgeInsets.all(14),
                        border: InputBorder.none,
                        hintText: 'How do you feel right now?',
                        hintStyle: kBodyTextStyle.copyWith(
                          color: kTextLight,
                          fontSize: 14,
                        ),
                      ),
                    ),
                  ),
                  const SizedBox(height: 40),
                  Row(
                    children: [
                      Expanded(
                        child: OutlinedButton(
                          onPressed: () => Navigator.pop(context),
                          style: OutlinedButton.styleFrom(
                            foregroundColor: kPrimaryBrown,
                            side: const BorderSide(color: kPrimaryBrown),
                          ),
                          child: const Text('Cancel'),
                        ),
                      ),
                      const SizedBox(width: 16),
                      Expanded(
                        child: ElevatedButton(
                          onPressed: provider.selectedMoodIndex != null || provider.note.isNotEmpty
                              ? provider.save
                              : null,
                          style: ElevatedButton.styleFrom(
                            backgroundColor: kPrimaryButtonColor,
                            foregroundColor: Colors.white,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          child: const Text('Save'),
                        ),
                      ),
                    ],
                  ),
                ] else ...[
                  SizedBox(
                    height: MediaQuery.of(context).size.height * 0.6,
                    child: Center(
                      child: Column(
                        mainAxisSize: MainAxisSize.min,
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Text(
                            'Successfully Saved!',
                            style: kBodyTextStyle.copyWith(
                              fontSize: 18,
                              fontWeight: FontWeight.w600,
                              color: kPrimaryBrown,
                            ),
                          ),
                          const SizedBox(height: 32),
                          ElevatedButton(
                            style: ElevatedButton.styleFrom(
                              backgroundColor: kPrimaryBrown,
                              foregroundColor: Colors.white,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(10),
                              ),
                              padding: const EdgeInsets.symmetric(horizontal: 40, vertical: 12),
                            ),
                            onPressed: provider.edit,
                            child: const Text('Edit'),
                          ),
                        ],
                      ),
                    ),
                  ),
                ]
              ],
            ),
          );
        },
      ),
    );
  }
}