import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const ContentUri = "/utsukushii.json"
export const DevModeCheck = "/dev_mode/check"
export const DevModeCmdUri = "/dev_mode/cmd"
