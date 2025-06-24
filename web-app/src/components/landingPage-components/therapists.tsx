"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { ChevronLeft, ChevronRight } from "lucide-react";
import {
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselNext,
  CarouselPrevious,
} from "../ui/carousel";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";

interface Therapist {
  id: number;
  name: string;
  initials: string;
  specialty: string;
  description?: string;
}

const therapists: Therapist[] = [
  {
    id: 1,
    name: "Dr. Sarah Johnson",
    initials: "SJ",
    specialty: "Clinical Psychologist",
  },
  {
    id: 2,
    name: "Dr. Michael Chen",
    initials: "MC",
    specialty: "Marriage & Family Therapist",
  },
  {
    id: 3,
    name: "Dr. Emily Rodriguez",
    initials: "ER",
    specialty: "Trauma Specialist",
  },
  {
    id: 4,
    name: "Dr. David Kim",
    initials: "DK",
    specialty: "Cognitive Behavioral Therapist",
  },
  {
    id: 5,
    name: "Dr. Lisa Thompson",
    initials: "LT",
    specialty: "Child & Adolescent Therapist",
  },
  {
    id: 6,
    name: "Dr. James Wilson",
    initials: "JW",
    specialty: "Addiction Counselor",
  },
];

const Therapists = () => {
  return (
    <section className="py-16 md:px-16 px-5">
      <div className="text-center mb-12">
        <h2 className="text-4xl font-bold text-[#4D4D4D] mb-6">
          Meet Our Expert Therapists
        </h2>
        <p className="text-gray-500 max-w-3xl mx-auto leading-relaxed">
          Our team of licensed professionals is here to support you on your
          journey to mental wellness. Each therapist brings unique expertise and
          compassionate care to help you achieve your goals.
        </p>
      </div>
      <Carousel
        opts={{
          align: "start",
        }}
        className="w-full relative"
      >
        <CarouselContent>
          {therapists.map((therapist, index) => (
            <CarouselItem key={index} className="md:basis-1/2 lg:basis-1/4">
              <Card>
                <CardContent className="flex flex-col aspect-square justify-center items-center p-6 space-y-10">
                  <Avatar className="mb-4 w-32 h-32">
                    <AvatarImage src="https://github.com/shadcn.png" />
                    <AvatarFallback>{therapist.initials}</AvatarFallback>
                  </Avatar>
                  <h3 className="text-xl font-semibold text-gray-800 mb-2">
                    {therapist.name}
                  </h3>
                  <p className="text-gray-600 mb-6">{therapist.specialty}</p>
                  <Button className="bg-green-500 hover:bg-green-600 text-white px-8 py-2 rounded-md">
                    Book Session
                  </Button>
                </CardContent>
              </Card>
            </CarouselItem>
          ))}
        </CarouselContent>
        <div className="absolute -top-6 right-8 flex gap-2">
          <CarouselPrevious />
          <CarouselNext />
        </div>
      </Carousel>
    </section>
  );
};

export default Therapists;
