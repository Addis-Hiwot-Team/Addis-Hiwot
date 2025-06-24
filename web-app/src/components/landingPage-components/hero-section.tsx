import Image from "next/image";
import React from "react";
import { Card, CardContent } from "../ui/card";
import { Bot, GraduationCap, Users } from "lucide-react";
import Link from "next/link";

const HeroSection = () => {
  return (
    <div>
      <div className="bg-[#F5F7FA] flex flex-col md:flex-row px-4 md:px-12 lg:px-32 py-10 md:py-16 lg:py-20 items-center md:space-x-2 space-y-8 md:space-y-0">
        <div className="space-y-4 text-left w-full md:w-2/3">
          <p className="font-bold text-4xl md:text-5xl lg:text-6xl text-[#4CAF4F]">
            <span className="text-[#4D4D4D]">Empowering Your</span> Recovery
            Journey.
          </p>
          <p className="text-[#4D4D4D] text-sm md:text-base lg:text-lg">
            Our comprehensive recovery app provides personalized support, expert
            guidance, and a compassionate community to help you achieve lasting
            sobriety.
          </p>
          <Link href="/api/auth" className="bg-[#4CAF4F] text-white px-4 py-2 rounded w-full sm:w-auto">
            Register
          </Link>
        </div>
        <div className="relative w-full h-80 md:w-1/3 flex justify-center items-center">
            <Image
              src="/images/hero.svg"
              alt="hero"
              fill
              className="object-cover rounded-tr-4xl rounded-bl-4xl "
            />
        </div>
      </div>
      <div className="py-10 md:py-16 px-2 md:px-4">
        <div className="w-full mx-auto text-center">
          <div className="mb-10 md:mb-16">
            <h2 className="text-2xl sm:text-3xl md:text-4xl font-bold text-[#4D4D4D] mb-4">
              Comprehensive Support For
              <br />
              your Recovery.
            </h2>
            <p className="text-gray-600 mx-auto text-base md:text-lg">
              Our recovery app offers a wide range of features to support you
              throughout your addiction recovery journey.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <Card className="border-0 bg-transparent">
              <CardContent className="p-6 text-center">
                <div className="w-16 h-16 mx-auto mb-6 bg-green-100 rounded-full flex items-center justify-center">
                  <Users className="w-8 h-8 text-green-600" />
                </div>
                <h3 className="text-xl md:text-2xl font-bold text-gray-800 mb-4">
                  Get Therapist
                  <br />
                  Assistant
                </h3>
                <p className="text-gray-600 leading-relaxed text-base md:text-lg">
                  Our membership management software provides full automation of
                  membership renewals and payments
                </p>
              </CardContent>
            </Card>

            <Card className="border-0 bg-transparent">
              <CardContent className="p-6 text-center">
                <div className="w-16 h-16 mx-auto mb-6 bg-green-100 rounded-full flex items-center justify-center">
                  <GraduationCap className="w-8 h-8 text-green-600" />
                </div>
                <h3 className="text-xl md:text-2xl font-bold text-gray-800 mb-4">
                  Education
                </h3>
                <p className="text-gray-600 leading-relaxed text-base md:text-lg">
                  Our membership management software provides full automation of
                  membership renewals and payments
                </p>
              </CardContent>
            </Card>

            <Card className="border-0 bg-transparent">
              <CardContent className="p-6 text-center">
                <div className="w-16 h-16 mx-auto mb-6 bg-green-100 rounded-full flex items-center justify-center">
                  <Bot className="w-8 h-8 text-green-600" />
                </div>
                <h3 className="text-xl md:text-2xl font-bold text-gray-800 mb-4">
                  AI Assistant
                </h3>
                <p className="text-gray-600 leading-relaxed text-base md:text-lg">
                  Our membership management software provides full automation of
                  membership renewals and payments
                </p>
              </CardContent>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

export default HeroSection;
