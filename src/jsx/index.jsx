'use strict';

var request = require('superagent');
var React = require('react');
var Bootstrap = require('react-bootstrap');
var ButtonToolbar = Bootstrap.ButtonToolbar;
var ButtonGroup = Bootstrap.ButtonGroup;
var Button = Bootstrap.Button;
var Modal = Bootstrap.Modal;
var ModalTrigger = Bootstrap.ModalTrigger;
var Glyphicon = Bootstrap.Glyphicon
var LinkButton = require('./_link-button.jsx');

var Page = React.createClass({
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
