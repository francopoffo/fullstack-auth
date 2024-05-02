import RegisterForm from "@/components/register/RegisterForm";

const Register = () => {
  return (
    <section className="bg-slate-200 mx-auto h-screen w-full flex items-center justify-center">
      <div className="bg-white w-1/2 mx-auto rounded-md h-fit p-4">
        <RegisterForm />
      </div>
    </section>
  );
};

export default Register;
