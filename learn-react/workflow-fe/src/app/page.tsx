import ObjectRender from "./component/ObjectRender";
import {Quiz1} from "./component/Quiz1";
import Quiz2 from "./component/Quiz2";
import Quiz3 from "./component/Quiz3";
import Quiz4 from "./component/Quiz4";

export default function Home() {
  const name = "turk"
  const age = 20
  const alive = true 
  return (
    <>
    <div>
      <Quiz1/>
    </div>
    <div>
      <Quiz2/>
    </div>
    <div>
      <Quiz3/>
    </div>
    <div>
      <Quiz4/>
    </div>
    </>
  );
}
