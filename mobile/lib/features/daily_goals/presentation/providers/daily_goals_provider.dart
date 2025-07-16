import 'package:flutter/material.dart';

class DailyGoalsProvider with ChangeNotifier {
  final List<Map<String, dynamic>> _goals = [
    {
      'title': 'Drink 8 Glasses of Water',
      'desc': 'Track your daily water intake.',
      'date': 'JULY 1, 2024',
      'tag': 'Daily',
      'done': false,
      'statusColor': Color(0xFF4CD964),
    },
    {
      'title': 'Read for 30 Minutes',
      'desc': 'Spend at least 30 minutes reading a book.',
      'date': 'JULY 1, 2024',
      'tag': 'Daily',
      'done': false,
      'statusColor': Color(0xFFFFD600),
    },
    {
      'title': 'Exercise',
      'desc': 'Complete a 20-minute workout.',
      'date': 'JULY 1, 2024',
      'tag': 'Daily',
      'done': true,
      'statusColor': Color(0xFF4CD964),
    },
    {
      'title': 'Meditate',
      'desc': 'Practice mindfulness for 10 minutes.',
      'date': 'JULY 1, 2024',
      'tag': 'Daily',
      'done': false,
      'statusColor': Color(0xFFFFD600),
    },
  ];

  List<Map<String, dynamic>> get goals => List.unmodifiable(_goals);

  void addGoal(Map<String, dynamic> goal) {
    _goals.add(goal);
    notifyListeners();
  }

  void removeGoal(int index) {
    _goals.removeAt(index);
    notifyListeners();
  }

  void toggleDone(int index) {
    _goals[index]['done'] = !_goals[index]['done'];
    notifyListeners();
  }

  void clearGoals() {
    _goals.clear();
    notifyListeners();
  }
} 