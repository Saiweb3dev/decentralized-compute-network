import { Button } from "@/components/ui/button"

export default function Hero() {
  return (
    <section className="w-full h-screen flex justify-center items-center py-20 bg-black">
      <div className="container  mx-auto text-center">
        <h1 className="text-4xl md:text-6xl  text-white font-bold mb-6">Decentralized Distributed Compute</h1>
        <p className="text-cyan-300 text-xl mb-8 max-w-2xl mx-auto">
          Boost productivity and simplify your business processes with StreamLine's powerful SaaS platform.
        </p>
        <Button size="lg" variant="secondary">
          Start Free Trial
        </Button>
      </div>
    </section>
  )
}

