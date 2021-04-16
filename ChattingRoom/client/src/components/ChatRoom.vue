<template>
    <el-container style="height: 100%;">
        <el-header height="120px" style="background-color: peachpuff; display: grid">
            <el-row :gutter="20" type="flex" justify="space-around" align="middle">
                <el-col :span="6" :offset="0" style="font-size: xx-large;">Chatting Room</el-col>
                <el-col :span="6" :offset="12" style="font-size: xx-large">Nickname：{{ nickname }}</el-col>
            </el-row>
        </el-header>
        <el-main style="border: 1px; border-radius: 2px;">
            <el-input
                    type="textarea"
                    readonly
                    v-model="messages"
                    style="height: 100%;"
            ></el-input>
        </el-main>
        <el-footer height="240px" style="border: 1px; border-radius: 4px; box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1); padding: 5px;">
            <el-input
                    type="textarea"
                    :rows="5"
                    placeholder="请输入消息内容"
                    v-model="message"
                    style="border-radius: 2px;"
            ></el-input>
            <el-row :gutter="20" type="flex" justify="end" align="middle" style="padding: 10px;">
                <el-button type="primary" size="medium" @click="send()" v-if="nickname !== ''">发送</el-button>
                <el-button type="primary" size="medium" @click="disconnect()" v-if="nickname !== ''">退出聊天</el-button>
                <el-button type="primary" size="medium" @click="join()" v-if="nickname === ''">加入聊天</el-button>
            </el-row>
        </el-footer>
    </el-container>
</template>

<script>
    import {Notification} from 'element-ui';
    let base_url = 'ws://127.0.0.1:8081/ws';

    export default {
        name: 'ChatRoom',
        data(){
            return {
                nickname: '',
                message: '',
                messages: '',
                socket: null,
            }
        },
        methods: {
            send() {
                let info = this.nickname + '：' + this.message + '\n';
                let data = {
                    user: this.nickname,
                    message: this.message
                };
                this.socket.send(JSON.stringify(data));
                this.messages += info;
                this.message = '';
            },
            join() {
                this.$prompt('请输入昵称', '提示', {
                    confirmButtonText: '确定',
                    inputPlaceholder: '请输入昵称',
                    inputErrorMessage: '昵称不能为空',
                    inputValidator: function ($event) {
                    return $event.length > 0
                    }
                }).then(({ value }) => {
                    this.nickname = value;
                //    发起 websocket 连接
                    this.createWebSocket();
                }).catch(() => {
                    console.log('取消输入')
                })
            },
            disconnect() {
                if (this.socket !== undefined && this.socket !== null) {
                    this.socket.close();
                }
                this.socket = null;
                this.nickname = '';
            },
            leaving($event) {
                console.log($event);
                console.log(this.socket);
                if (this.socket !== undefined && this.socket !== null) {
                    this.socket.close();
                }
            },
            createWebSocket() {
                if (this.socket === null) {
                    this.socket = new WebSocket(base_url)
                }

                this.socket.onopen = (event) => {
                    console.log(event);
                    console.log('connected');
                    Notification.success('连接已建立');
                //    初次建立连接，发送昵称
                    let data = {user: this.nickname};
                    this.socket.send(JSON.stringify(data));
                };

                this.socket.onmessage = (event) => {
                    console.log(event);
                    let data = JSON.parse(event.data)
                    let info = data.user + '：' + data.message + '\n';
                    Notification.info(info);
                    //    消息内容处理
                    this.messages += info
                };

                this.socket.onerror = (event) => {
                    console.log(event);
                    console.log('error');
                    Notification.error('服务端运行发生异常');
                };

                this.socket.onclose = (event) => {
                    console.log(event);
                    console.log('close');
                    Notification.info('服务端关闭连接');
                    this.nickname = '';
                    this.socket = null;
                };
            },
        },
        created() {
            window.addEventListener('beforeunload', this.leaving);
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
>>> .el-textarea__inner {
        height: 100%;
    }
</style>
