<template>
    <div class="card">
        <div class="card-body">
            <div class="row">
                <div class="col-3">
                    <img
                        class="img-fluid"
                        src="https://s2.loli.net/2022/01/26/egdCNVunqMRoiUY.png"
                        alt=""
                    />
                </div>
                <div class="col-9">
                    <div class="username">{{ user.username }}</div>
                    <div class="fans">粉丝：{{ user.followerCount }}</div>
                    <button
                        @click="follow"
                        v-if="!user.isFollowed"
                        type="button"
                        class="btn btn-secondary btn-sm"
                    >
                        +关注
                    </button>
                    <button
                        @click="unfollow"
                        v-else
                        type="button"
                        class="btn btn-secondary btn-sm"
                    >
                        取消关注
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { computed } from "vue";

export default {
    name: "UserInfo",
    props: {
        user: {
            type: Object,
            required: true,
        },
    },
    setup(props, context) {
        let fullName = computed(
            () => props.user.lastName + " " + props.user.firstName
        );
        const follow = () => {
            context.emit("followEvent");
            console.log("follow");
        };
        const unfollow = () => {
            context.emit("unfollowEvent");
            console.log("unfollow");
        };

        return {
            fullName,
            follow,
            unfollow,
        };
    },
};
</script>

<style scoped>
img {
    border-radius: 50%;
}

.username {
    font-weight: bold;
}

.fans {
    font-size: 12px;
    color: gray;
}

button {
    padding: 2px 4px;
    font-size: 12px;
}

.img-field {
    display: flex;
    flex-direction: column;
    justify-content: center;
}
</style>
