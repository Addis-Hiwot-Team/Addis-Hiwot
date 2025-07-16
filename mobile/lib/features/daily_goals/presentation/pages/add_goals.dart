// lib/screens/daily_goals/add_goal_screen.dart

import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
import 'package:mobile/core/constants/constants.dart';
import 'package:provider/provider.dart';
import '../providers/daily_goals_provider.dart';

class AddGoalScreen extends StatefulWidget {
  const AddGoalScreen({super.key});

  @override
  State<AddGoalScreen> createState() => _AddGoalScreenState();
}

class _AddGoalScreenState extends State<AddGoalScreen> {
  String? _selectedFrequency = 'Daily';
  final _titleController = TextEditingController();
  final _descController = TextEditingController();
  final _deadlineController = TextEditingController(text: '12/12/2025');

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Color(0xFFD7CCC8),
      appBar: PreferredSize(
        preferredSize: const Size.fromHeight(110), // <-- increase height
        child: SafeArea(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 8),
            child: Column(
              mainAxisSize: MainAxisSize.min, // <-- add this
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                // Back arrow at the top left
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
                // Title and icon
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
      body: Center(
        child: Container(
          width: 320,
          margin: const EdgeInsets.only(top: 16, bottom: 24),
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 24),
          decoration: BoxDecoration(
            color: kWhite,
            borderRadius: BorderRadius.circular(24),
          ),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                'Add New Goal',
                style: GoogleFonts.poppins(fontSize: 18, fontWeight: FontWeight.bold, color: kPrimaryBrown),
              ),
              const SizedBox(height: 18),
              _buildTextField(label: 'Title', hint: 'Typing |', controller: _titleController),
              const SizedBox(height: 18),
              _buildTextField(label: 'Description', hint: 'Typing |', controller: _descController, maxLines: 3),
              const SizedBox(height: 18),
              _buildTextField(
                label: 'Deadline',
                hint: '12/12/2025',
                controller: _deadlineController,
                readOnly: true,
                suffixIcon: Icons.calendar_today_outlined,
                onTap: () async {
                  DateTime? picked = await showDatePicker(
                    context: context,
                    initialDate: DateTime.tryParse(_deadlineController.text.split('/').reversed.join('-')) ?? DateTime.now(),
                    firstDate: DateTime(2000),
                    lastDate: DateTime(2100),
                  );
                  if (picked != null) {
                    setState(() {
                      _deadlineController.text = "${picked.month.toString().padLeft(2, '0')}/${picked.day.toString().padLeft(2, '0')}/${picked.year}";
                    });
                  }
                },
              ),
              const SizedBox(height: 18),
              _buildDropdownField(
                label: 'Frequency',
                value: _selectedFrequency,
                items: const ['Daily', 'Weekly', 'Monthly', 'Yearly'],
                onChanged: (val) => setState(() => _selectedFrequency = val),
              ),
              const SizedBox(height: 28),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    backgroundColor: kPrimaryBrown,
                    foregroundColor: Colors.white,
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(14),
                    ),
                    padding: const EdgeInsets.symmetric(vertical: 12),
                    elevation: 0,
                  ),
                  onPressed: () {
                    Provider.of<DailyGoalsProvider>(context, listen: false).addGoal({
                      'title': _titleController.text,
                      'desc': _descController.text,
                      'date': _deadlineController.text,
                      'tag': _selectedFrequency,
                      'done': false,
                      'statusColor': Color(0xFFFFD600),
                    });
                    Navigator.pop(context);
                  },
                  child: const Text('Save', style: TextStyle(fontSize: 16)),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  Widget _buildTextField({
    required String label,
    required String hint,
    TextEditingController? controller,
    int maxLines = 1,
    bool readOnly = false,
    IconData? suffixIcon,
    VoidCallback? onTap,
  }) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(label, style: GoogleFonts.poppins(fontWeight: FontWeight.w600, color: kPrimaryBrown)),
        const SizedBox(height: 8),
        TextField(
          controller: controller,
          maxLines: maxLines,
          readOnly: readOnly,
          onTap: onTap,
          style: GoogleFonts.poppins(color: kPrimaryBrown),
          decoration: InputDecoration(
            hintText: hint,
            hintStyle: GoogleFonts.poppins(color: kTextLight),
            filled: true,
            fillColor: kGreyColor,
            contentPadding: const EdgeInsets.symmetric(horizontal: 12, vertical: 12),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: const BorderSide(color: kTextLight),
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: const BorderSide(color: kPrimaryBrown, width: 1.5),
            ),
            suffixIcon: suffixIcon != null ? Icon(suffixIcon, color: kTextLight) : null,
          ),
        ),
      ],
    );
  }

  Widget _buildDropdownField({
    required String label,
    required String? value,
    required List<String> items,
    required ValueChanged<String?> onChanged,
  }) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(label, style: GoogleFonts.poppins(fontWeight: FontWeight.w600, color: kPrimaryBrown)),
        const SizedBox(height: 8),
        DropdownButtonFormField<String>(
          value: value,
          decoration: InputDecoration(
            filled: true,
            fillColor: kGreyColor,
            contentPadding: const EdgeInsets.symmetric(horizontal: 12, vertical: 12),
            enabledBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: const BorderSide(color: kTextLight),
            ),
            focusedBorder: OutlineInputBorder(
              borderRadius: BorderRadius.circular(10),
              borderSide: const BorderSide(color: kPrimaryBrown, width: 1.5),
            ),
          ),
          icon: const Icon(Icons.keyboard_arrow_down_rounded, color: kTextLight),
          style: GoogleFonts.poppins(color: kPrimaryBrown),
          items: items.map((String val) {
            return DropdownMenuItem<String>(
              value: val,
              child: Text(val),
            );
          }).toList(),
          onChanged: onChanged,
        ),
      ],
    );
  }
}