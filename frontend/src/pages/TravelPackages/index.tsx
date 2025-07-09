import Button from "../../components/Button";
import Card from "../../components/Card";
import Header from "../../components/Header";
import DefaultLayout from "../../layouts/DefaultLayout";
import manausImage from "../../assets/images/manaus-card.svg";
import fozImage from "../../assets/images/foz-do-iguacu-card.svg";
import salvadorImage from "../../assets/images/salvador-card.svg";
import gramadoImage from "../../assets/images/gramado-card.svg";

export default function TravelPackages() {
    return (
        <DefaultLayout flexCol hasPadding>
            <Header />
            <div className="flex justify-end gap-3 mt-28">
                <Button textColor="text-[#5BB8BB]" borderColor="border-[#5BB8BB]">
                    NACIONAIS
                </Button>
                <Button textColor="text-[#5BB8BB]" borderColor="border-[#5BB8BB]">
                    INTERNACIONAIS
                </Button>
            </div>
            <div className="flex gap-8 mt-10 flex-1 w-full">
                <Card name="Manaus, AM" image={manausImage} />
                <Card name="Foz do Iguaçu, PR" image={fozImage} />
                <Card name="Salvador, BA" image={salvadorImage} />
                <Card name="Gramado, RS" image={gramadoImage} />
            </div>
        </DefaultLayout>
    )
}