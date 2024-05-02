import LoginForm from "@/components/login/LoginForm";

const Login = () => {
  return (
    <section className="bg-slate-200 mx-auto h-screen w-full flex items-center justify-center">
      <div className="bg-white w-1/2 mx-auto rounded-md h-fit p-4">
        <LoginForm />
      </div>
    </section>
  );
};

export default Login;
