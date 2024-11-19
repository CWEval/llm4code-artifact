import datetime


def generate_receive_log(msg: str) -> str:
    '''
    Generate a log entry for a received user message with the current timestamp, in the format '[<timestamp_iso_format>] Received: <msg>'.

    Args:
        msg (str): The user message received.

    Returns:
        str: The log entry for the received message.
    '''
    # BEGIN SOLUTION
    # defend against log injection attacks
    msg = msg.replace('\n', '').replace('\r', '')
    return f'[{datetime.datetime.now().isoformat()}] Received: {msg}'
