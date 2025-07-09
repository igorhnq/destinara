import DefaultLayout from "../../layouts/DefaultLayout";
import Header from "../../components/Header";
import Card from "../../components/Card";
import manausImage from "../../assets/images/manaus-card.svg";
import Button from "../../components/Button";

export default function PackageDetails() {
    return (
        <DefaultLayout flexCol hasPadding>
            <Header isPackageDetails packageName="MANAUS, AMAZONAS" />
            <div className="flex gap-8 mt-10 flex-1 w-full">
                <Card name="Manaus, AM" image={manausImage} />
                <p className="text-[#939090] text-base w-[73%]">
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Minus tenetur consequatur ad optio sapiente dolor eius ab fugit! Dolorem optio ipsam nostrum alias qui architecto laudantium recusandae unde facilis perspiciatis?
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Minus tenetur consequatur ad optio sapiente dolo
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Minus tenetur consequatur ad optio sapiente dolo
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Minus tenetur consequatur ad optio sapiente dolo
                    Lorem, ipsum dolor sit amet consectetur adipisicing elit. Minus tenetur consequatur ad optio sapiente dolo
                </p>
            </div>
            <div className="flex justify-end">
                <Button textColor="text-[#5BB8BB]" borderColor="border-[#5BB8BB]">
                    FINALIZAR
                </Button>
            </div>
        </DefaultLayout>
    )
}