# plugins

The directory stores plugins for the app
Please contribute new plugins here.

## Specification

Just send http post request to http://{domain_name}/apps/:id/cmd

*body*

```json
{
    "name": "{command_name}",
    "accessToken": "{access_token_of_an_app}"
}
```

## Plugins implemented

[twitchchat](./twitchchat)
