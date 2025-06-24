import { MoveRight } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import React from "react";

const Header = () => {
  return (
    <div className="flex justify-between px-5 md:px-20 items-center space-x-5 sm:space-x-0">
      <div className="relative min-h-16 min-w-16 h-20 w-20">
        <Image src="/images/logo.svg" alt="logo" fill className="" />
      </div>
      <div className="flex space-x-2 md:space-x-4 items-center md:font-semibold text-[#4D4D4D] text-xs md:text-base">
        <div>Home</div>
        <div>Features</div>
        <div>Community</div>
        <Link href="/api/auth" className="hidden bg-[#4CAF4F] text-white px-2 sm:px-4 py-2 rounded sm:flex space-x-2 items-center text-xs sm:text-sm whitespace-nowrap">
          <span> Register Now</span> <MoveRight size={16} />
        </Link>
      </div>
    </div>
  );
};

export default Header;
