"use server";
import {getOauhProviders, getUserInfo, getUserTokens, userInfo} from "@/api/getUserInfos";
import ProfilePage from "@/components/pages/profile/ProfilePage";

export default async function Profile() {
    try {
        const data: userInfo = await getUserInfo();
        data.usersProviders = await getUserTokens();
        data.possibleProviders = await getOauhProviders();
        return <ProfilePage {...data} />;
    } catch (err) {
        const data: userInfo = {id: -1, first_name: "", last_name: "", email: "", password: "", usersProviders: [], possibleProviders: []};
        console.log(err)
        return <ProfilePage {...data} />
    }
}
