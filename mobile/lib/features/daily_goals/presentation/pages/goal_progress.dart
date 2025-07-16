// lib/screens/goal_progress_screen.dart

import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';
 // <-- THIS IS THE MISSING LINE TO ADD

class GoalProgressScreen extends StatelessWidget {
  const GoalProgressScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(Icons.arrow_back_ios_new),
          onPressed: () => Navigator.of(context).popUntil((route) => route.isFirst),
        ),
        title: Text('Progress', style: GoogleFonts.poppins(fontWeight: FontWeight.bold)),
        centerTitle: true,
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            const Spacer(),
            SizedBox(
              width: 200,
              height: 200,
              child: CustomPaint(
                painter: StreakProgressPainter(
                  purpleProgress: 0.8, // 80%
                  pinkProgress: 0.5,   // 50%
                ),
              ),
            ),
            const SizedBox(height: 40),
            Text(
              '5 Days Streak !',
              textAlign: TextAlign.center,
              style: GoogleFonts.poppins(
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
            ),
            const Spacer(),
            Row(
              children: [
                Expanded(
                  child: OutlinedButton(
                    onPressed: () => Navigator.of(context).popUntil((route) => route.isFirst),
                    child: const Text('Return'),
                  ),
                ),
                const SizedBox(width: 16),
                Expanded(
                  child: ElevatedButton(
                    onPressed: () {},
                    child: const Text('Save'),
                  ),
                ),
              ],
            ),
            const SizedBox(height: 20),
          ],
        ),
      ),
    );
  }
}



// lib/widgets/streak_progress_painter.dar
// This is the DEFINITION of the painter.
// It extends CustomPainter, which means it promises
// to have a `paint` method and a `shouldRepaint` method.
class StreakProgressPainter extends CustomPainter {
  final double purpleProgress;
  final double pinkProgress;

  // The constructor gets the data it needs to draw.
  StreakProgressPainter({
    required this.purpleProgress,
    required this.pinkProgress,
  });

  // This is the core method. Flutter calls this when it's time to draw.
  // It gives you a `canvas` (your drawing surface) and a `size` (the area to draw in).
  @override
  void paint(Canvas canvas, Size size) {
    // ... all the drawing logic is here ...
    // e.g., creating Paint objects and calling canvas.drawArc()
  }

  // This tells Flutter whether it needs to redraw if the painter's properties change.
  @override
  bool shouldRepaint(covariant CustomPainter oldDelegate) {
    return true; // We set it to true to keep it simple.
  }
}