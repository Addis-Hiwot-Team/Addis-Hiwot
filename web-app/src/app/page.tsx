import About from "@/components/landingPage-components/about";
import Footer from "@/components/landingPage-components/footer";
import Header from "@/components/landingPage-components/header";
import HeroSection from "@/components/landingPage-components/hero-section";
import Quotes from "@/components/landingPage-components/quotes";
import Therapists from "@/components/landingPage-components/therapists";

export default function Home() {
  return (
    <div className="flex flex-col min-h-screen">
      <Header />
      <HeroSection />
      <About />
      <Therapists />
      <Quotes />
      <Footer />
    </div>
  );
}
