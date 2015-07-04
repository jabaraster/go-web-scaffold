'use strict';

var React = require('react');
var Bootstrap = require('react-bootstrap');
var Input = Bootstrap.Input;
var Popover = Bootstrap.Popover;
var OverlayTrigger = Bootstrap.OverlayTrigger;
var InputField = require('./_input-field.jsx');

var Page = React.createClass({
    handleValueChange: function(e) {
        console.log(e);
    },
    render: function() {
        return (
            <div className="Page container">
                <h1>This is login page!</h1>
                <InputField label="ユーザID"
                            required={true}
                            placeholder="メールアドレス形式"
                            maxCharCount={100}
                            type="text"
                            initialValue="userId"
                            descriptor="userId"
                            onValueChange={this.handleValueChange}
                />
            </div>
        );
    }
});

React.render(
    <Page />,
    document.getElementById('page')
);
