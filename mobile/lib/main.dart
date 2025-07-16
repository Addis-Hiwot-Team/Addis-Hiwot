import 'package:flutter/material.dart';
import 'package:device_preview/device_preview.dart';
import 'package:mobile/features/Dashboard/screen/dash_board.dart';
import 'package:mobile/features/auth/presentation/login.dart';
import 'package:provider/provider.dart';

// Pages
import 'package:mobile/features/landing/landing_page.dart';
import 'package:mobile/features/auth/presentation/sign_up.dart';
import 'package:mobile/features/auth/presentation/password_reset/forget_pass.dart';
import 'package:mobile/features/auth/presentation/password_reset/verify_email.dart';
import 'package:mobile/features/auth/presentation/password_reset/reset_page.dart';
import 'package:mobile/features/auth/presentation/password_reset/success.dart';
import 'package:mobile/features/daily_goals/presentation/pages/daily_goals_list_screen.dart';
import 'package:mobile/features/Daily_checkin/presentation/pages/daily_checkin_screen.dart';

// Auth Layer Imports
import 'package:mobile/features/auth/presentation/providers/auth_provider.dart';
import 'package:mobile/features/auth/data/repository/user_repository_impl.dart';
import 'package:mobile/features/auth/data/data_sources/local_datasource/auth_local_datasource.dart';
import 'package:mobile/features/auth/data/data_sources/remote_datasource/auth_remote_datasource.dart';
import 'package:mobile/core/utils/network_info.dart';
import 'package:mobile/features/auth/domain/usecase/login_usecase.dart';
import 'package:mobile/features/auth/domain/usecase/signup_usecase.dart';
import 'package:mobile/features/auth/domain/usecase/logout_usecase.dart';
import 'package:mobile/features/auth/domain/usecase/getme_usecase.dart';

import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:http/http.dart' as http;
import 'package:connectivity_plus/connectivity_plus.dart';

void main() {
  // Initialize dependencies
  final localDataSource = AuthLocalDataSourceImpl(storage: const FlutterSecureStorage());
  final remoteDataSource = AuthRemoteDatasourceImpl(client: http.Client());
  final networkInfo = NetworkInfo(Connectivity());

  final userRepository = UserRepositoryImpl(
    local: localDataSource,
    remote: remoteDataSource,
    networkInfo: networkInfo,
  );

  final loginUseCase = LoginUseCase(userRepository);
  final signupUseCase = SignupUseCase(userRepository);
  final logoutUseCase = LogoutUseCase(userRepository);
  final getMeUseCase = GetMeUseCase(userRepository);

  runApp(
    DevicePreview(
      enabled: true,
      builder: (context) => MultiProvider(
        providers: [
          ChangeNotifierProvider(
            create: (_) => AuthProvider(
              loginUseCase: loginUseCase,
              signupUseCase: signupUseCase,
              logoutUseCase: logoutUseCase,
              getMeUseCase: getMeUseCase,
            ),
          ),
          // Add other providers here when needed
          // ChangeNotifierProvider(create: (_) => AnotherProvider()),
        ],
        child: const MyApp(),
      ),
    ),
  );
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Addis Hiwot',
      debugShowCheckedModeBanner: false,
      useInheritedMediaQuery: true, // Required by DevicePreview
      locale: DevicePreview.locale(context),
      builder: DevicePreview.appBuilder,
      initialRoute: '/dashboard',
      routes: {
        '/': (context) => const LandingPage(),
        '/login': (context) => const LoginPage(),
        '/signup': (context) => const SignUpPage(),
        '/forgot_password': (context) => const ForgotPasswordPage(),
        '/verify_email': (context) => const VerifyEmailPage(),
        '/reset_password': (context) => const ResetPasswordPage(),
        '/success': (context) => const SuccessPage(),
        '/dashboard': (context) => const DashboardScreen(),
        '/daily_goals': (context) => const DailyGoalsListScreen(),
        '/daily_checkin': (context) =>  DailyCheckInScreen(),
      },
    );
  }
}
