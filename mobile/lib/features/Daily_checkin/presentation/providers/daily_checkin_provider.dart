import 'package:flutter/material.dart';

class DailyCheckInProvider with ChangeNotifier {
  int? _selectedMoodIndex;
  String _note = '';
  bool _isSaved = false;

  int? get selectedMoodIndex => _selectedMoodIndex;
  String get note => _note;
  bool get isSaved => _isSaved;

  void selectMood(int index) {
    _selectedMoodIndex = index;
    notifyListeners();
  }

  void setNote(String value) {
    _note = value;
    notifyListeners();
  }

  void save() {
    _isSaved = true;
    notifyListeners();
  }

  void edit() {
    _isSaved = false;
    notifyListeners();
  }

  void reset() {
    _selectedMoodIndex = null;
    _note = '';
    _isSaved = false;
    notifyListeners();
  }
} 