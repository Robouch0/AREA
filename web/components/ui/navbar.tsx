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
    const img = prop.profilePicture || "/default-avatar.png";
    return (
        <div className="flex flex-row">
            <div className="p-4 ml-auto">
                <button
                    className="focus:border-slate-500 focus:border-4 rounded-3xl focus:outline-none focus:p-1"
                    onClick={() => (redirect("/services/profile"))}
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
