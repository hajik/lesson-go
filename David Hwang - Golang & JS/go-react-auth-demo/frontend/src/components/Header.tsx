import React, { SyntheticEvent } from 'react'
import {Navbar, Nav, Container} from 'react-bootstrap'

interface Props{
    firstName: string
    setFirstName: (firstName: string) => void
}

const Header = ({firstName, setFirstName}: Props) => {

    const logoutHandler = async(e: SyntheticEvent) => {
        e.preventDefault();

        // interact with the backend
        await fetch('http://localhost:8081/api/logout', {
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
        })

        setFirstName('')
    }

    return (
        <Navbar bg="dark" variant="dark" expand="lg" collapseOnSelect>
        <Container>
            <Navbar.Brand href="/">Go React Auth</Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
                {firstName ? (
                    <Nav className="me-auto">
                        <Nav.Link onClick={logoutHandler}>Logout</Nav.Link>
                    </Nav>
                ) : (
                    <Nav className="me-auto">
                        <Nav.Link href="/">Home</Nav.Link>
                        <Nav.Link href="/signup">Sign Up</Nav.Link>
                        <Nav.Link href="/signin">Signin</Nav.Link>
                    </Nav>
                )}
            </Navbar.Collapse>
        </Container>
        </Navbar>
    )
}

export default Header
