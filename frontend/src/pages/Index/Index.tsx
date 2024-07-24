import coffePattern from "@/assets/coffee-pattern.jpg";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

function Index() {
  return (
    <>
      <div className="h-16 flex items-center border-b-border border-2 bg-red p-4">
        <img
          src="/public/favicon.svg"
          className="h-full"
          alt="Coffe cup icon"
        />
        <h1 className="text-lg">Coffee Assistant</h1>
        <div className="flex-grow"></div>
        <div>Help</div>
      </div>
      <div
        className="bg-contain bg-repeat h-96 flex items-center justify-center"
        style={{ backgroundImage: `url("${coffePattern}")` }}
      >
        <Card className="w-96">
          <CardHeader>
            <CardTitle>Want to make coffee?</CardTitle>
          </CardHeader>
          <CardContent>
            <form>
              <div className="grid w-full intems-center gap-4">
                <div className="flex space-x-1.5">
                  <Input
                    id="coffemaker"
                    placeholder="Type your coffemaker, ex: Chemex"
                  />
                  <Button>Start</Button>
                </div>
              </div>
            </form>
          </CardContent>
        </Card>
      </div>
    </>
  );
}

export default Index;
