const React = require('react');
const Bootstrap = require('react-bootstrap');
const Button = Bootstrap.Button;
const Glyphicon = Bootstrap.Glyphicon;

const LinkButton = React.createClass({
    propTypes: {
        href: React.PropTypes.string.isRequired,
        text: React.PropTypes.string,
        glyph: React.PropTypes.string,
        style: React.PropTypes.string
    },
    handleClick: function() {
        location.href = this.props.href;
    },
    render: function() {
        const glyph = this.props.glyph ? (<Glyphicon glyph={this.props.glyph} />) : (<span/>);
        return (
            <Button bsStyle={this.props.style} onClick={this.handleClick}>
                {glyph}
                {this.props.text}
            </Button>
        );
    }
});

module.exports = LinkButton;
