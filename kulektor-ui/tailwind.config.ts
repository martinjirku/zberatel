import type { Config } from "tailwindcss";
import { mtConfig } from "@material-tailwind/react";

const config: Config = {
  content: [
    "./src/**/*.{jsx,tsx,ts,js}",
    "./src/main.tsx",
    "./index.html",
    "./node_modules/@material-tailwind/react/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
    screens: {
      sm: "480px",
      md: "768px",
      lg: "976px",
      xl: "1440px",
    },
  },
  plugins: [
    mtConfig({
      radius: "1px",
    }),
  ],
};

export default config;
