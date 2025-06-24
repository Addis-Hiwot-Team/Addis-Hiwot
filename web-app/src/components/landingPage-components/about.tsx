import { BookHeart, CalendarClock, Club, Users } from "lucide-react";
import Image from "next/image";
import React from "react";

const About = () => {
  return (
    <div>
      <div className="sm:flex px-5 md:px-24 py-20 items-center sm:space-x-10 md:space-x-20 space-y-4 sm:space-y-0">
        <div className="sm:w-1/3 flex justify-center items-center">
          <div className="relative w-full sm:w-70 h-70">
            <Image
              src="/images/about.svg"
              alt="hero"
              fill
              className="object-cover rounded-2xl"
            />
          </div>
        </div>
        <div className="space-y-4 text-left sm:w-2/3">
          <p className="font-bold text-3xl text-[#4D4D4D]">
            Where Healing and Hope Begin
          </p>
          <div className="h-36 overflow-hidden w-full">
            <p className="text-gray-500 text ">
              At our center, we believe that true healing begins with feeling
              seen, heard, and deeply cared for. We don’t just provide services
              — we walk beside you, offering patience, understanding, and
              genuine compassion every step of the way. Here, you are welcomed
              like family and supported with kindness and respect for your
              unique story. Whether you’re seeking help for mental health
              challenges or freedom from addiction, we offer a safe space where
              you can breathe, grow, and heal at your own pace. We know that
              asking for help takes courage, and we honor that by meeting you
              with warmth, encouragement, and unwavering support. No one should
              face life’s struggles alone. We’re here to remind you that hope is
              real, change is possible, and you have the strength to write a new
              chapter. Take things one day at a time — and know that every small
              step forward is worth celebrating.
            </p>
          </div>
          <button className="bg-[#4CAF4F] text-white px-2 py-1 rounded">
            Learn More
          </button>
        </div>
      </div>
      <div className="bg-[#F5F7FA] sm:flex px-5 sm:px-10 py-20 items-center sm:space-x-5">
        <div className="space-y-4 text-left w-full lg:w-1/2 p-5">
          <p className="font-bold text-5xl text-[#4CAF4F]">
            <span className="text-[#4D4D4D]">
              Helping a local Helping our community{" "}
            </span>
            find hope and healing
          </p>
          <p className="text-[#4D4D4D]">
            Our comprehensive recovery app provides personalized support, expert
            guidance, and a compassionate community to help you achieve lasting
            sobriety.
          </p>
        </div>
        <div className="w-full lg:w-1/2 grid grid-cols-2 gap-10 p-5 ">
          <div className="flex items-center space-x-5 ">
            <div>
              <Users />
            </div>
            <div className="">
              <p className="font-bold text-[#4D4D4D] text-2xl">2,245,341</p>
              <p className="text-gray-400 text-sm">Members</p>
            </div>
          </div>
          <div className="flex items-center space-x-5">
            <div>
              <Club />
            </div>
            <div>
              <p className="font-bold text-[#4D4D4D] text-2xl">1,245,341</p>
              <p className="text-gray-400 text-sm">Club</p>
            </div>
          </div>
          <div className="flex items-center space-x-5">
            <div>
              <CalendarClock />
            </div>

            <div>
              <p className="font-bold text-[#4D4D4D] text-2xl">1,245,341</p>
              <p className="text-gray-400 text-sm">Event</p>
            </div>
          </div>
          <div className="flex items-center space-x-5">
            <div>
              <BookHeart />
            </div>
            <div>
              <p className="font-bold text-[#4D4D4D] text-2xl">1,245,341</p>
              <p className="text-gray-400 text-sm">Therapist</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default About;
