import { useState } from "react";
import { Button } from "@/components/ui/button";
import { DevModeCmdUri } from "@/lib/utils";

export interface DevModeRunProps {
  reload: React.EventHandler<any>;
}

export function DevModeRun(props: DevModeRunProps) {
  const [isLoading, setIsLoading] = useState(false);

  const handleClick = async () => {
    setIsLoading(true);
    try {
      await fetch(DevModeCmdUri);
    } catch (error) {
      console.error("Error on run tests:", error);
    } finally {
      setIsLoading(false);
      props.reload(null)
    }
  };

  return (
    <Button onClick={handleClick} style={{ justifyContent: "flex-end" }}>
      <div className={isLoading ? "animate-bounce" : ""}>
        <TestTubeIcon />
      </div>
    </Button>
  );
}

function TestTubeIcon(props: React.SVGProps<SVGSVGElement>) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      height="20px"
      viewBox="0 -960 960 960"
      width="100px"
      fill="#434343"
    >
      <path d="M200-120q-51 0-72.5-45.5T138-250l222-270v-240h-40q-17 0-28.5-11.5T280-800q0-17 11.5-28.5T320-840h320q17 0 28.5 11.5T680-800q0 17-11.5 28.5T640-760h-40v240l222 270q32 39 10.5 84.5T760-120H200Zm80-120h400L544-400H416L280-240Zm-80 40h560L520-492v-268h-80v268L200-200Zm280-280Z" />

      <path d="M750-0v-560l440 280Zm80-280Zm0 134 210-134-210-134v268Z" />
    </svg>
  );
}
