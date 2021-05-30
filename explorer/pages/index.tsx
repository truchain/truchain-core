import React from "react";
import { motion } from "framer-motion";

export default function IndexPage() {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 h-screen w-screen overflow-y-hidden">
      <div className="h-screen hidden md:block">
        <img
          src="https://source.unsplash.com/random"
          className="object-cover h-screen w-full"
        />
      </div>
      <div className="h-screen flex flex-col items-center justify-center">
        <h1 className="text-5xl font-bold tracking-wide">TRUCHAIN</h1>
        <div className="mt-12 flex">
          <motion.div animate={{ scale: 1.1 }} transition={{ duration: 0.9 }}>
            <div className="rounded-lg shadow bg-gray-800 p-4 flex justify-center items-center m-4">
              <span className="text-3xl font-bold text-white">20</span>
            </div>
          </motion.div>
          <motion.div animate={{ scale: 1.1 }} transition={{ duration: 0.9 }}>
            <div className="rounded-lg shadow bg-gray-800 p-4 flex justify-center items-center m-4">
              <span className="text-3xl font-bold text-white">06</span>
            </div>
          </motion.div>
          <motion.div animate={{ scale: 1.1 }} transition={{ duration: 0.9 }}>
            <div className="rounded-lg shadow bg-gray-800 p-4 flex justify-center items-center m-4">
              <span className="text-3xl font-bold text-white">49</span>
            </div>
          </motion.div>
          <motion.div animate={{ scale: 1.1 }} transition={{ duration: 0.9 }}>
            <div className="rounded-lg shadow bg-gray-800 p-4 flex justify-center items-center m-4">
              <span className="text-3xl font-bold text-white">34</span>
            </div>
          </motion.div>
        </div>
      </div>
    </div>
  );
}
