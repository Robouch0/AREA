'use client';
import Image from 'next/image'
import {redirect} from "next/navigation";
// import aboutJson from "../../public/about.json";

export interface User {
    profilePicture?: string;
    imgWidth?: number;
    imgHeight?: number;
}


export default function Navbar({prop}: {prop:User}) {
    const img: string = prop.profilePicture || "/default-avatar.png";
    return (
        <div className="flex flex-row">
            <div className="p-4 ml-auto">
                <button
                    className="focus-visible:border-slate-500 focus-visible:border-4 rounded-3xl focus-visible:outline-none focus-visible:p-1"
                    onClick={(): never => (redirect("/services/profile"))}
                >
                    <Image
                        className="border-4 rounded-b-full rounded-t-full border-black hover:animate-pulse hover:ease-in "
                        alt="profile picture"
                        src={img}
                        width={prop.imgWidth}
                        height={prop.imgHeight}
                    />
                </button>
            </div>
        </div>
    );
}
