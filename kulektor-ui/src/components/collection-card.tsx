import {
  Card,
  CardBody,
  CardFooter,
  CardHeader,
} from "@material-tailwind/react";
import { PropsWithChildren } from "react";

interface Props {
  title: string;
}
export default function ColectionCard({
  children,
  title,
}: PropsWithChildren<Props>) {
  return (
    <Card>
      <CardHeader>{title}</CardHeader>
      <CardBody>{children}</CardBody>
      <CardFooter></CardFooter>
    </Card>
  );
}
