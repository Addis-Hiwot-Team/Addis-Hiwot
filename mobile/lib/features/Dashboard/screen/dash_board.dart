import 'package:flutter/material.dart';
import 'package:fl_chart/fl_chart.dart';
import 'package:mobile/core/constants/constants.dart';
import 'package:provider/provider.dart';
import '../provider/dashboard_provider.dart';


class DashboardScreen extends StatelessWidget {
  const DashboardScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => DashboardProvider(),
      child: Consumer<DashboardProvider>(
        builder: (context, provider, _) {
          return Scaffold(
            body: SafeArea(
              child: SingleChildScrollView(
                child: Padding(
                  padding: const EdgeInsets.all(20.0),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      _buildHeader(),
                      const SizedBox(height: 20),
                      _buildSearchBar(),
                      const SizedBox(height: 25),
                      _buildQuoteCard(),
                      const SizedBox(height: 25),
                      _buildActionButtons(),
                      const SizedBox(height: 25),
                      _buildMoodTracker(provider),
                      const SizedBox(height: 25),
                      _buildGoalsCompleted(),
                      const SizedBox(height: 25),
                      _buildHabits(),
                      const SizedBox(height: 25),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [_buildTalkWithAddisbot(),],)
                      
                    ],
                  ),
                ),
              ),
            ),
            bottomNavigationBar: _buildBottomNavBar(context, provider),
          );
        },
      ),
    );
  }

  Widget _buildHeader() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        Row(
          children: [
            const CircleAvatar(
              radius: 24,
              backgroundImage: AssetImage('assets/images/avatar.png'),
            ),
            const SizedBox(width: 12),
            Text(
              'Kidist',
              style: TextStyle(
                fontSize: 22,
                fontWeight: FontWeight.bold,
                color: kDarkBrownColor,
              ),
            ),
          ],
        ),
        Container(
          padding: const EdgeInsets.all(8),
          decoration: BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.circular(12),
            boxShadow: [
              BoxShadow(
                color: Colors.grey.withOpacity(0.2),
                spreadRadius: 1,
                blurRadius: 5,
              )
            ],
          ),
          child: const Icon(Icons.notifications_outlined, color: kDarkBrownColor),
        ),
      ],
    );
  }

  Widget _buildSearchBar() {
    return TextField(
      decoration: InputDecoration(
        hintText: 'Search...',
        hintStyle: TextStyle(color: Colors.grey.shade500),
        prefixIcon: Icon(Icons.search, color: Colors.grey.shade500),
        filled: true,
        fillColor: Colors.white,
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(12),
          borderSide: BorderSide.none,
        ),
        contentPadding: const EdgeInsets.symmetric(vertical: 0),
      ),
    );
  }

 Widget _buildQuoteCard() {
  return 
  ClipRRect(
    child:
  Container(
    height: 200,
    width: double.infinity,
    padding: const EdgeInsets.all(24),
    decoration: BoxDecoration(
      color: const Color(0xFFB79C86),
      borderRadius: BorderRadius.circular(20),
    ),
    child: const Stack(
      clipBehavior: Clip.none, // Allows content to overflow
      children: [
        // Decorative circles in the background
        Positioned(
          right: 5, // Move slightly off the right edge
          top: 15,  // Move down from the top
          child: CircleAvatar(
            radius: 70, // Slightly larger for more overlap
            backgroundColor: Color(0xFF946B4A),
          ),
        ),
        Positioned(
          left: -5,
          bottom: -34, // Move the bottom edge below the card
          child: CircleAvatar(
            radius: 30, // Smaller circle
            backgroundColor: Color(0xFF946B4A),
          ),
        ),
        // Greeting at the top left
        Positioned(
          left: 0,
          top: 0,
          child: Text(
            'Hello Kidist!',
            style: TextStyle(color: Colors.white70, fontSize: 16),
          ),
        ),
        // Center the quote and writer vertically but align left horizontally
        Align(
          alignment: Alignment.centerLeft,
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                '"Small steps matter."',
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 26,
                  fontWeight: FontWeight.bold,
                ),
                textAlign: TextAlign.left,
              ),
              SizedBox(height: 8),
              Text(
                '- By Someone',
                style: TextStyle(color: Colors.white70, fontSize: 14),
                textAlign: TextAlign.left,
              ),
            ],
          ),
        ),
      ],
    ),
  ) );
}

  Widget _buildActionButtons() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        _actionButton(Icons.checklist_rtl, 'Check-in', (context) {
          Navigator.pushNamed(context, '/daily_checkin');
        }),
        _actionButton(Icons.track_changes, 'Goal', (context) {
          Navigator.pushNamed(context, '/daily_goals');
        }),
        _actionButton(Icons.sync_alt, 'Habits', (context) {}),
        _actionButton(Icons.people_outline, 'Community', (context) {}),
      ],
    );
  }

  Widget _actionButton(IconData icon, String label, void Function(BuildContext) onTap) {
    return Builder(
      builder: (context) => GestureDetector(
        onTap: () => onTap(context),
        child: Container(
          padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 16),
          decoration: BoxDecoration(
            color: Colors.white,
            borderRadius: BorderRadius.circular(15),
            border: Border.all(color: Colors.grey.shade300),
          ),
          child: Column(
            children: [
              Icon(icon, color: kDarkBrownColor, size: 28),
              const SizedBox(height: 4),
              Text(label, style: const TextStyle(color: kDarkBrownColor, fontSize: 12)),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildMoodTracker(DashboardProvider provider) {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            'Mood Tracker',
            style: TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.bold,
              color: kDarkBrownColor,
            ),
          ),
          const SizedBox(height: 16),
          _buildTimeframeSelector(provider),
          const SizedBox(height: 16),
          _buildDateNavigator(),
          const SizedBox(height: 20),
          _buildMoodChart(),
        ],
      ),
    );
  }

  Widget _buildTimeframeSelector(DashboardProvider provider) {
    return Row(
      children: [
        _timeframeButton('Weekly', provider),
        const SizedBox(width: 8),
        _timeframeButton('Monthly', provider),
        const SizedBox(width: 8),
        _timeframeButton('Yearly', provider),
      ],
    );
  }

  Widget _timeframeButton(String title, DashboardProvider provider) {
    bool isSelected = provider.moodTrackerTimeframe == title;
    return GestureDetector(
      onTap: () {
        provider.setMoodTrackerTimeframe(title);
      },
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 8, horizontal: 16),
        decoration: BoxDecoration(
          color: isSelected ? kBrownColor : Colors.transparent,
          borderRadius: BorderRadius.circular(20),
          border: Border.all(
            color: isSelected ? Colors.transparent : Colors.grey.shade300,
          ),
        ),
        child: Text(
          title,
          style: TextStyle(
            color: isSelected ? Colors.white : kDarkBrownColor,
            fontWeight: FontWeight.w500,
          ),
        ),
      ),
    );
  }

  Widget _buildDateNavigator() {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceBetween,
      children: [
        const Icon(Icons.arrow_back_ios, size: 16, color: kDarkBrownColor),
        const Text(
          'June 08 - June 14, 2025',
          style: TextStyle(color: kDarkBrownColor, fontWeight: FontWeight.w500),
        ),
        const Icon(Icons.arrow_forward_ios, size: 16, color: kDarkBrownColor),
      ],
    );
  }

  Widget _buildMoodChart() {
    final List<double> weeklyMoods = [4, 2, 5, 4.5, 3, 1.5, 4];

    return SizedBox(
      height: 150,
      child: BarChart(
        BarChartData(
          alignment: BarChartAlignment.spaceAround,
          maxY: 5,
          minY: 0,
          barTouchData: BarTouchData(enabled: false),
          titlesData: FlTitlesData(
            leftTitles: AxisTitles(sideTitles: _moodYAxis()),
            rightTitles: const AxisTitles(),
            topTitles: const AxisTitles(),
            bottomTitles: AxisTitles(sideTitles: _moodXAxis()),
          ),
          gridData: const FlGridData(show: false),
          borderData: FlBorderData(show: false),
          barGroups: List.generate(weeklyMoods.length, (index) {
            return BarChartGroupData(
              x: index,
              barRods: [
                BarChartRodData(
                  toY: weeklyMoods[index],
                  color: index == 2 ? kHappyYellowColor : kGreyColor,
                  width: 20,
                  borderRadius: const BorderRadius.all(Radius.circular(6)),
                ),
              ],
            );
          }),
        ),
      ),
    );
  }

  SideTitles _moodYAxis() {
    const style = TextStyle(fontSize: 20);
    return SideTitles(
      showTitles: true,
      interval: 1,
      getTitlesWidget: (value, meta) {
        switch (value.toInt()) {
          case 1:
            return const Text('😵', style: style);
          case 2:
            return const Text('😠', style: style);
          case 3:
            return const Text('😐', style: style);
          case 4:
            return const Text('😊', style: style);
          case 5:
            return const Text('😄', style: style);
          default:
            return const SizedBox.shrink();
        }
      },
      reservedSize: 40,
    );
  }
  
  SideTitles _moodXAxis() {
    const style = TextStyle(color: kDarkBrownColor, fontSize: 12);
    return SideTitles(
      showTitles: true,
      getTitlesWidget: (value, meta) {
        String text = '';
        switch (value.toInt()) {
          case 0: text = '08'; break;
          case 1: text = '09'; break;
          case 2: text = '10'; break;
          case 3: text = '11'; break;
          case 4: text = '12'; break;
          case 5: text = '13'; break;
          case 6: text = '14'; break;
        }
        return Padding(
          padding: const EdgeInsets.only(top: 8.0),
          child: Text(text, style: style),
        );
      },
      reservedSize: 30,
    );
  }


  Widget _buildGoalsCompleted() {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            'Goals Completed',
            style: TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.bold,
              color: kDarkBrownColor,
            ),
          ),
          const SizedBox(height: 8),
          const Text(
            '2 out of 5 completed',
            style: TextStyle(color: Colors.grey, fontSize: 14),
          ),
          const SizedBox(height: 12),
          Row(
            children: [
              Expanded(
                child: ClipRRect(
                  borderRadius: BorderRadius.circular(10),
                  child: const LinearProgressIndicator(
                    value: 0.45,
                    minHeight: 12,
                    backgroundColor: kGreyColor,
                    valueColor: AlwaysStoppedAnimation<Color>(kOrangeColor),
                  ),
                ),
              ),
              const SizedBox(width: 12),
              const Text(
                '45%',
                style: TextStyle(
                  color: kDarkBrownColor,
                  fontWeight: FontWeight.bold,
                  fontSize: 16,
                ),
              ),
            ],
          ),
        ],
      ),
    );
  }

  Widget _buildHabits() {
    return Container(
      padding: const EdgeInsets.all(16),
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: BorderRadius.circular(20),
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const Text(
            'Habits',
            style: TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.bold,
              color: kDarkBrownColor,
            ),
          ),
          const SizedBox(height: 16),
          _habitItem('Sleep at 10pm', true),
          const Divider(height: 24),
          _habitItem('Drink 2 liter of water', false),
        ],
      ),
    );
  }

  Widget _habitItem(String title, bool isStreak) {
    return Row(
      children: [
        const Text('• ', style: TextStyle(color: kDarkBrownColor)),
        Text(title, style: const TextStyle(color: kDarkBrownColor, fontSize: 16)),
        const Spacer(),
        if (isStreak)
          const Column(
            children: [
              Text('🔥 4', style: TextStyle(color: Colors.orange, fontWeight: FontWeight.bold)),
              Text('days streak', style: TextStyle(color: Colors.grey, fontSize: 12)),
            ],
          )
        else
          const Column(
            children: [
              Icon(Icons.check_circle, color: Colors.green),
              Text('Done', style: TextStyle(color: Colors.grey, fontSize: 12)),
            ],
          ),
      ],
    );
  }

  Widget _buildTalkWithAddisbot() {
    return ElevatedButton(
      onPressed: () {},
      style: ElevatedButton.styleFrom(
        backgroundColor: Colors.white,
        foregroundColor: kDarkBrownColor,
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(15)),
        padding: const EdgeInsets.symmetric(vertical: 12, horizontal: 20),
        elevation: 2,
      ),
      child: Row(
        mainAxisSize: MainAxisSize.min,
        children: [
          const Icon(Icons.smart_toy, size: 40), // Chatbot icon
          const SizedBox(width: 12),
          const Text(
            'Talk with Addisbot',
            style: TextStyle(fontSize: 16, fontWeight: FontWeight.w500),
          ),
        ],
      ),
    );
  }

  Widget _buildBottomNavBar(BuildContext context, DashboardProvider provider) {
    return Container(
      decoration: BoxDecoration(
        color: Colors.white,
        borderRadius: const BorderRadius.only(
          topLeft: Radius.circular(25),
          topRight: Radius.circular(25),
        ),
        boxShadow: [
          BoxShadow(
            color: Colors.grey.withOpacity(0.2),
            spreadRadius: 1,
            blurRadius: 10,
          )
        ],
      ),
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 20.0, vertical: 15.0),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: [
            _bottomNavItem(Icons.home, 'Home', 0, provider),
            _bottomNavItem(Icons.library_books_outlined, 'Resource', 1, provider),
            _bottomNavItem(Icons.chat_bubble_outline, 'Messages', 2, provider),
            _bottomNavItem(Icons.person_outline, 'Profile', 3, provider),
          ],
        ),
      ),
    );
  }

  Widget _bottomNavItem(IconData icon, String label, int index, DashboardProvider provider) {
    bool isSelected = provider.selectedTab == index;
    return GestureDetector(
      onTap: () {
        provider.setTab(index);
      },
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Icon(
            icon,
            color: isSelected ? kBrownColor : Colors.grey.shade400,
            size: 28,
          ),
          const SizedBox(height: 4),
          Text(
            label,
            style: TextStyle(
              color: isSelected ? kBrownColor : Colors.grey.shade400,
              fontSize: 12,
              fontWeight: isSelected ? FontWeight.bold : FontWeight.normal,
            ),
          ),
        ],
      ),
    );
  }
}