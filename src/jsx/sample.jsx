'use strict';

const request = require('superagent');
const React = require('react');
const Bootstrap = require('react-bootstrap');
const ButtonToolbar = Bootstrap.ButtonToolbar;
const ButtonGroup = Bootstrap.ButtonGroup;
const Button = Bootstrap.Button;
const Modal = Bootstrap.Modal;
const ModalTrigger = Bootstrap.ModalTrigger;
const Glyphicon = Bootstrap.Glyphicon
const LinkButton = require('./component/link-button.jsx');
const InputField = require('./component/input-field.jsx');
const FormInputMixin = require('./mixin/form-input.js');

// Modalはコンポーネントで包まないとクローズボタンが動作しなくなる. 不可解・・・
const ModalSample = React.createClass({
    handle: function() {
        if (this.props.onClose() === false) {
            return;
        }
        this.props.onRequestHide(); // これを呼ぶとモーダルが消える. 不可解・・・
    },
    render: function() {
        return (
           <Modal {...this.props}
                  title={`モーダルダイアログ`}
                  animation={true}
                  backdrop={false}
           >
               <div className="modal-body">
                   本文
               </div>
               <div className="modal-footer">
                   <Button bsStyle="primary" onClick={this.handle}>Close</Button>
               </div>
           </Modal>
        );
    }
});

const Page = React.createClass({
    mixins: [FormInputMixin],
    afterHandleValueChange: function() {
        console.log(arguments);
    },
    handleClick: function() {
        console.log(this.hasError());
    },
    handleModalClose: function() {
        return confirm('閉じる？');
    },
    handleLogout: function() {
        location.href = "/logout";
    },
    render: function() {
        // onRequestHideハンドラを指定しても動かない. 不可解・・・
        const modal = (
            <ModalSample onClose={this.handleModalClose} />
        );
        return (
            <div className="Page container">
                <form>
                    <InputField label="ユーザID"
                                required={true}
                                placeholder="メールアドレス形式"
                                maxCharCount={100}
                                type="text"
                                initialValue=""
                                descriptor="userId"
                                onValueChange={this.handleValueChange}
                    />
                    <InputField label="パスワード"
                                required={true}
                                placeholder="メールアドレス形式"
                                maxCharCount={100}
                                type="text"
                                initialValue=""
                                descriptor="password"
                                onValueChange={this.handleValueChange}
                    />
                    <Button bsStyle="success" onClick={this.handleClick}><Glyphicon glyph="ok" /></Button>
                </form>
                <h1>Bootstrap Components !</h1>
                <ButtonToolbar>
                    <Button bsStyle="default">Default</Button>
                    <Button bsStyle="primary">Primary</Button>
                    <Button bsStyle="success">Success</Button>
                    <Button bsStyle="warning">Warning</Button>
                    <Button bsStyle="danger">Danger</Button>
                </ButtonToolbar>
                <hr />
                <ButtonGroup>
                </ButtonGroup>
                <hr />
                <ModalTrigger modal={modal}>
                    <Button bsStyle="primary"><Glyphicon glyph="search" /> {`Show Dialog`}</Button>
                </ModalTrigger>
                <hr />
                <LinkButton style="default" text="Reload" glyph="refresh" href="/" />
                <LinkButton style="primary" text="Logout" glyph="log-out" href="/logout" />
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);
