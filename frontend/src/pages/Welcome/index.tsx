import Button from "../../components/Button";
import { Link } from "react-router-dom";

export default function Welcome() {
    return (
        <>
            <div className="flex flex-col items-center justify-center gap-10 h-screen bg-[#7DDADD]">
                <div className="flex items-center justify-center gap-5">
                    <div className="flex items-center justify-center border-white border-2 w-[275px] h-[275px] ">logo</div>
                    <div className="flex items-center justify-center text-center flex-col gap-2">
                        <h1 className="text-white text-6xl font-extrabold">Bem-vindo</h1>
                        <p className="text-white text-xl font-semibold ">Faça seu login ou cadastre-se</p>
                    </div>
                </div>
                <div className="flex gap-5">
                    <Link to="/login" className="block">
                        <Button>
                            LOGIN
                        </Button>
                    </Link>
                    <Link to="/cadastro" className="block">
                        <Button>
                            CADASTRO
                        </Button>
                    </Link>
                </div>
            </div>
        </>
    )
}