'use strict';

const React = require('react');
const Bootstrap = require('react-bootstrap');
const ButtonToolbar = Bootstrap.ButtonToolbar;
const ButtonGroup = Bootstrap.ButtonGroup;
const Button = Bootstrap.Button;
const Modal = Bootstrap.Modal;
const ModalTrigger = Bootstrap.ModalTrigger;
const Glyphicon = Bootstrap.Glyphicon
const LinkButton = require('./component/link-button.jsx');

const Page = React.createClass({
    render: function() {
        return (
            <div className="Page container">
                <ButtonGroup>
                    <LinkButton style="primary" text="Reload" glyph="refresh" href="/" />
                    <LinkButton style="primary" text="Logout" glyph="log-out" href="/logout" />
                </ButtonGroup>
                <ButtonGroup>
                    <LinkButton style="default" text="サンプルページ" href="/page/sample/" />
                </ButtonGroup>
                <ButtonGroup>
                    <LinkButton style="success" text="ユーザマスタメンテ" href="/page/app-user/" />
                </ButtonGroup>
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);
