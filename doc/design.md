# Design

## Workflow 

### Normal workflow

[End user] calls an API to send a message which contains a `Service ID` to call-of-duty.
[System] routes the message to the service which matches the `Service ID`.
[System] sends the message to the active escalation policy under the service.
[System] starts to call the policy layers on the policy chain, passing the message on.
[System] goes into a calendar entry of a policy layer.
[System] finds the target user to notify.
[System] sends the message upon the user's notification rules.
[Notified User] receives the message.
[Notified User] acknowledges the message.
[System] receives the acknowledge confirmation.
[System] stops passing the message on the policy chain.

### Policy chain loop

[System] sends the message upon the user's notification rules.
[Notified User] receives the message, but no ACK.
[System] times out waiting for the acknowledgement.
[System] calls the next layer on the chain.
[System] calls the loop policy (ONCE, LOOP) when it comes to the end of the chain.

## Data Model

### Message


