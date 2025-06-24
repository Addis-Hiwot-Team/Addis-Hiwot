import React from "react";
import { Card, CardContent } from "../ui/card";
import Image from "next/image";
import { ArrowRight } from "lucide-react";

const supportiveNotes = [
  {
    id: 1,
    quote: "It always seems impossible until it is done.",
    image: "/placeholder.svg?height=200&width=200",
    alt: "Inspirational portrait",
  },
  {
    id: 2,
    quote: "We may encounter many defeats but we must not be defeated.",
    image: "/placeholder.svg?height=200&width=200",
    alt: "Inspirational portrait",
  },
  {
    id: 3,
    quote:
      "Although the world is full of suffering, it is also full of the overcoming of it.",
    image: "/placeholder.svg?height=200&width=200",
    alt: "Inspirational portrait",
  },
];

const Quotes = () => {
  return (
    <div className="px-5 md:px-20">
      <div className="text-center mb-12">
        <h2 className="text-4xl font-bold text-[#4D4D4D] mb-6">
          Supportive Notes Section
        </h2>
        <p className="text-gray-500 max-w-3xl mx-auto leading-relaxed ">
          Until we gather stories from those we’ve helped, here are some gentle
          reminders and wise words to lift your spirit and remind you: healing
          is possible, and you’re not alone.
        </p>
      </div>
      <div className="md:flex justify-between items-center md:space-x-5 space-y-14 md:space-y-0">
        {supportiveNotes.map((note) => (
          <Card
            key={note.id}
            className="relative bg-white shadow-lg hover:shadow-xl transition-shadow duration-300 w-full p-0 overflow-visible rounded-2xl"
          >
            <CardContent className="p-0 min-h-80 flex flex-col justify-center rounded-2xl">
              <Image
                src="https://github.com/shadcn.png"
                alt="Quote Image"
                fill
                className="object-cover grayscale rounded-t-2xl"
              />

              <div className="absolute left-1/2 -translate-x-1/2 -bottom-10 p-8 w-[90%] bg-white flex flex-col items-center rounded-2xl shadow-md">
                <blockquote className="text-gray-700 text-lg min-h-[80px] flex items-center justify-center text-center">
                  "{note.quote}"
                </blockquote>

                <button className="inline-flex items-center text-green-600 hover:text-green-700 font-medium transition-colors duration-200 group mt-4">
                  Readmore
                  <ArrowRight className="ml-2 h-4 w-4 group-hover:translate-x-1 transition-transform duration-200" />
                </button>
              </div>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
};

export default Quotes;
