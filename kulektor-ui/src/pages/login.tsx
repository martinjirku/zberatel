import { Button, Input } from "@material-tailwind/react";

export default function Login() {
  return (
    <form>
      <div className="mb-4">
        <Input type="email" size="lg" className="w-full" required />
      </div>
      <div className="mb-6">
        <Input type="password" size="lg" className="w-full" required />
      </div>
      <Button className="w-full" color="primary" type="submit">
        Login
      </Button>
    </form>
  );
}
