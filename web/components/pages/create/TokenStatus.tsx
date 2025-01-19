import React from 'react';

export default function TokenStatus({isTokenActionPresent, isTokenReactionPresent, actionName, reactionName}: {
    isTokenActionPresent: boolean;
    isTokenReactionPresent: boolean;
    actionName: string;
    reactionName: string;
}) {
    if (isTokenActionPresent && isTokenReactionPresent)
        return null;

    return (
        <div className={"font-bold mt-4 text-xl "}>
            <p>You cannot create this area</p>
            <p>There is no account linked to AREA for the following services:</p>
            {!isTokenActionPresent && <p className={"font-bold mx-4"}> Action: {actionName}</p>}
            {!isTokenReactionPresent && <p className={"font-bold mx-4"}> Reaction: {reactionName}</p>}
        </div>
    );
}
