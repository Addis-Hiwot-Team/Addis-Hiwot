import 'package:flutter/material.dart';

class DashboardProvider with ChangeNotifier {
  int _selectedTab = 0;
  String _moodTrackerTimeframe = 'Weekly';

  int get selectedTab => _selectedTab;
  String get moodTrackerTimeframe => _moodTrackerTimeframe;

  void setTab(int index) {
    _selectedTab = index;
    notifyListeners();
  }

  void setMoodTrackerTimeframe(String timeframe) {
    _moodTrackerTimeframe = timeframe;
    notifyListeners();
  }

  // Add more dashboard-specific state and methods as needed
} 