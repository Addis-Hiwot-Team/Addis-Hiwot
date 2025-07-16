// lib/screens/daily_goals/daily_goals_list_screen.dart

import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:mobile/core/constants/constants.dart';
import 'package:mobile/features/daily_goals/presentation/pages/add_goals.dart';
import 'package:provider/provider.dart';
import '../providers/daily_goals_provider.dart';

class DailyGoalsListScreen extends StatefulWidget {
  const DailyGoalsListScreen({super.key});

  @override
  State<DailyGoalsListScreen> createState() => _DailyGoalsListScreenState();
}

class _DailyGoalsListScreenState extends State<DailyGoalsListScreen> {
  String _filter = 'All';

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => DailyGoalsProvider(),
      child: Consumer<DailyGoalsProvider>(
        builder: (context, provider, _) {
          final goals = provider.goals;
          final hasGoals = goals.isNotEmpty;
          return Scaffold(
            backgroundColor: kBackgroundColor,
            appBar: PreferredSize(
              preferredSize: const Size.fromHeight(110),
              child: SafeArea(
                child: Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8),
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Container(
                        decoration: const BoxDecoration(
                          color: kGreyColor,
                          shape: BoxShape.circle,
                        ),
                        child: IconButton(
                          icon: const Icon(Icons.arrow_back_ios_new, color: kPrimaryBrown),
                          onPressed: () => Navigator.pop(context),
                        ),
                      ),
                      const SizedBox(height: 16),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.start,
                        children: [
                          Text(
                            'Daily Goals',
                            style: GoogleFonts.poppins(
                              fontWeight: FontWeight.bold,
                              fontSize: 22,
                              color: kPrimaryBrown,
                            ),
                          ),
                          const SizedBox(width: 16),
                          const Icon(Icons.track_changes, color: kPrimaryBrown, size: 28),
                        ],
                      ),
                    ],
                  ),
                ),
              ),
            ),
            body: Padding(
              padding: const EdgeInsets.symmetric(horizontal: 16.0),
              child: hasGoals
                  ? _buildGoalsContent(goals)
                  : Center(
                      child: Column(
                        mainAxisAlignment: MainAxisAlignment.center,
                        mainAxisSize: MainAxisSize.max,
                        children: [
                          Text('No goals yet!', style: kTitleStyle, textAlign: TextAlign.center),
                          const SizedBox(height: 10),
                          Text(
                            'Set your first goal to begin making progress.',
                            textAlign: TextAlign.center,
                            style: kBodyTextStyle.copyWith(color: kTextLight, fontSize: 16),
                          ),
                        ],
                      ),
                    ),
            ),
            floatingActionButton: Padding(
              padding: const EdgeInsets.only(bottom: 12.0),
              child: SizedBox(
                width: 200,
                height: 48,
                child: ElevatedButton.icon(
                  style: ElevatedButton.styleFrom(
                    backgroundColor: kPrimaryBrown,
                    foregroundColor: Colors.white,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(18),
                    ),
                    elevation: 0,
                  ),
                  onPressed: () {
                    Navigator.push(context, MaterialPageRoute(builder: (context) => const AddGoalScreen()));
                  },
                  icon: const Icon(Icons.add, color: Colors.white),
                  label: Text('Create Goal', style: GoogleFonts.poppins(fontWeight: FontWeight.w600)),
                ),
              ),
            ),
            floatingActionButtonLocation: FloatingActionButtonLocation.centerFloat,
          );
        },
      ),
    );
  }

  // _buildEmptyState removed, logic now in body

  Widget _buildGoalsContent(List<Map<String, dynamic>> goals) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const SizedBox(height: 8),
        _buildFilterTabs(),
        const SizedBox(height: 18),
        Expanded(
          child: ListView.separated(
            padding: const EdgeInsets.only(bottom: 90),
            itemCount: goals.length,
            separatorBuilder: (context, index) => const SizedBox(height: 14),
            itemBuilder: (context, index) {
              final goal = goals[index];
              return _GoalListItem(
                title: goal['title'],
                description: goal['desc'],
                date: goal['date'],
                tag: goal['tag'],
                isDone: goal['done'],
                statusColor: goal['statusColor'],
              );
            },
          ),
        ),
      ],
    );
  }

  Widget _buildFilterTabs() {
    final tabs = ['All', 'Active', 'Done'];
    return Container(
      width: double.infinity, // Make it full width
      padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
      decoration: BoxDecoration(
        color: kPrimaryBrown, // brown
        borderRadius: BorderRadius.circular(16),
      ),
      child: Row(
        // Remove mainAxisSize: MainAxisSize.min,
        mainAxisAlignment: MainAxisAlignment.spaceEvenly, // Distribute tabs evenly
        children: List.generate(tabs.length * 2 - 1, (i) {
          if (i.isOdd) {
            // Divider
            return Container(
              width: 1,
              height: 24,
              color: Colors.white.withOpacity(0.5),
              margin: const EdgeInsets.symmetric(horizontal: 8),
            );
          }
          final index = i ~/ 2;
          final isSelected = _filter == tabs[index];
          return Expanded(
            child: GestureDetector(
              onTap: () => setState(() => _filter = tabs[index]),
              child: Container(
                alignment: Alignment.center,
                padding: const EdgeInsets.symmetric(horizontal: 0, vertical: 8),
                decoration: isSelected
                    ? BoxDecoration(
                        color: kGreyColor, // lighter pill
                        borderRadius: BorderRadius.circular(12),
                      )
                    : null,
                child: Text(
                  tabs[index],
                  style: TextStyle(
                    color: isSelected ? kPrimaryBrown : Colors.white,
                    fontWeight: FontWeight.w600,
                  ),
                ),
              ),
            ),
          );
        }),
      ),
    );
  }
}

class _GoalListItem extends StatelessWidget {
  final String title;
  final String description;
  final String date;
  final String tag;
  final bool isDone;
  final Color statusColor;

  const _GoalListItem({
    required this.title,
    required this.description,
    required this.date,
    required this.tag,
    required this.isDone,
    required this.statusColor,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 0),
      padding: const EdgeInsets.all(14),
      decoration: BoxDecoration(
        color: const Color(0xFFE4C4A9), // <-- set your desired color here
        borderRadius: BorderRadius.circular(16),
      ),
      child: Stack(
        children: [
          Row(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Status circle
              Padding(
                padding: const EdgeInsets.only(top: 6.0),
                child: Container(
                  width: 18,
                  height: 18,
                  decoration: BoxDecoration(
                    shape: BoxShape.circle,
                    color: Colors.white,
                    border: Border.all(
                      color: isDone ? kPrimaryBrown : kTextLight,
                      width: 2,
                    ),
                  ),
                  child: isDone
                      ? const Icon(Icons.check, size: 14, color: kPrimaryBrown)
                      : null,
                ),
              ),
              const SizedBox(width: 12),
              // Goal details
              Expanded(
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text(title, style: kBodyTextStyle.copyWith(fontWeight: FontWeight.w600, fontSize: 15, color: kPrimaryBrown)),
                    const SizedBox(height: 2),
                    Text(description, style: kBodyTextStyle.copyWith(color: kPrimaryBrown.withOpacity(0.6), fontSize: 13)),
                    const SizedBox(height: 10),
                    Text(date, style: kBodyTextStyle.copyWith(color: kPrimaryBrown.withOpacity(0.4), fontSize: 12)),
                  ],
                ),
              ),
              const SizedBox(width: 10),
              // Status dot at the top right
              Column(
                crossAxisAlignment: CrossAxisAlignment.end,
                children: [
                  Container(
                    width: 10,
                    height: 10,
                    decoration: BoxDecoration(
                      color: statusColor,
                      shape: BoxShape.circle,
                    ),
                  ),
                ],
              ),
            ],
          ),
          // Tag at the bottom right
          Positioned(
            right: 0,
            bottom: 0,
            child: Text(
              tag,
              style: GoogleFonts.poppins(
                fontSize: 12,
                color: kPrimaryBrown,
                fontWeight: FontWeight.w600,
              ),
            ),
          ),
        ],
      ),
    );
  }
}