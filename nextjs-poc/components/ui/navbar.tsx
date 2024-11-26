'use client';
import Image from 'next/image'
import {redirect} from "next/navigation";
export default function Navbar({prop} :{prop: any}) {
    let img;
    if (prop.profilePicture != undefined) {
        img = prop.profilePicture;
    } else {
        img = "/default-avatar.png";
    }

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
