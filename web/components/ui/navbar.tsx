'use client';
import Image from 'next/image'
import {redirect} from "next/navigation";
import aboutJson from "../../public/about.json";

export interface User {
    profilePicture?: string;
    age?: number;
} // unsure is this the right away pour du TS type ?

export default function Navbar({prop}: {prop:User}) {
    const img = prop.profilePicture || "/default-avatar.png";
    console.log(aboutJson.services);
    return (
        <div className="flex flex-row">
            <div className="p-4 ml-auto">
                <button onClick={() => (redirect("/services/profile"))}>
                    <Image
                        className="border-4 rounded-b-full rounded-t-full border-black"
                        alt="profile picture"
                        src={img}
                        width={60}
                        height={60}
                    />
                </button>
            </div>
        </div>
    );
}
