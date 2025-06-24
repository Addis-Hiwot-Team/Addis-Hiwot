import { Instagram, MessageCircle, Twitter, Youtube } from "lucide-react";
import Image from "next/image";
import React from "react";

const Footer = () => {
  return (
    <div className="mt-20">
      <div className="bg-[#F8F9FA] rounded-b-xl pt-10 pb-12 px-4 text-center">
        <h2 className="text-4xl font-bold text-[#4D4D4D] mb-4">
          Ready to Start Your Journey?
        </h2>
        <p className="text-gray-600 max-w-2xl mx-auto mb-6">
          Take the first step towards better mental health. Our therapists are
          here to provide personalized care tailored to your unique needs and
          goals.
        </p>
        <button className="bg-green-500 hover:bg-green-600 text-white font-semibold px-8 py-3 rounded transition-colors">
          Schedule Consultation
        </button>
      </div>
      <footer className="bg-gray-900 text-gray-300 pt-12 pb-6 px-4">
        <div className="max-w-7xl mx-auto flex flex-col md:flex-row md:justify-between gap-8">
          <div className="flex flex-col gap-4 md:w-1/3">
            <div className="flex items-center gap-2">
              <Image
                src="/images/logo.svg"
                alt="logo"
                width={100}
                height={100}
              />
            </div>
            <div className="text-sm text-gray-400">
              Copyright © 2020 Landify UI Kit.
              <br />
              All rights reserved
            </div>
            <div className="flex gap-4 mt-2">
              <a href="#" aria-label="Instagram" className="hover:text-white">
                <Instagram />
              </a>
              <a href="#" aria-label="Dribbble" className="hover:text-white">
                <MessageCircle />
              </a>
              <a href="#" aria-label="Twitter" className="hover:text-white">
                <Twitter />
              </a>
              <a href="#" aria-label="YouTube" className="hover:text-white">
                <Youtube />
              </a>
            </div>
          </div>
          <div className="flex flex-1 justify-between gap-8">
            <div>
              <h3 className="font-semibold text-white mb-3">Company</h3>
              <ul className="space-y-2 text-sm">
                <li>
                  <a href="#" className="hover:text-white">
                    About us
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Blog
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Contact us
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Pricing
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Testimonials
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h3 className="font-semibold text-white mb-3">Support</h3>
              <ul className="space-y-2 text-sm">
                <li>
                  <a href="#" className="hover:text-white">
                    Help center
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Terms of service
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Legal
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Privacy policy
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-white">
                    Status
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="md:w-1/3">
            <h3 className="font-semibold text-white mb-3">Stay up to date</h3>
            <form className="flex items-center gap-2">
              <input
                type="email"
                placeholder="Your email address"
                className="rounded bg-gray-800 border border-gray-700 px-3 py-2 text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-green-400"
              />
              <button
                type="submit"
                className="p-2 rounded bg-green-500 hover:bg-green-600 text-white"
                aria-label="Subscribe"
              >
                <svg
                  width="18"
                  height="18"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth="2"
                    d="M5 12h14M12 5l7 7-7 7"
                  />
                </svg>
              </button>
            </form>
          </div>
        </div>
      </footer>
    </div>
  );
};

export default Footer;
