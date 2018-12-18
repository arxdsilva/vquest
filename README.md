# vQuest

vQuest is a simple MMORPG made with golang.

# TODO
- [ ] Only allow conn attempt when user/pass are filled
-     [ ] If user or pass not filled, 'illuminate' that box
- [ ] Deactivate client on `err`
- [ ] Deactivate client on `disconnection`
- [ ] Characters screen
-     [ ] New character button on CS
- [ ] BUG: Fix `selectedField` == 1 on start
- [ ] Refactor UI system
- [ ] New character screen
- [ ] Solve encoding problem with pass character '●'
- [ ] Add Connecting screen ?
- [ ] Do not use user infos as global value
- [ ] Store password values as hash. This will prevent memory manipulation to get the user pass

# Done
- [x] Simple `user connected` communication screen
- [x] implement blinking char on `selectedField`
- [x] Fix blink Login button on `enter` key
- [x] allign text input with blinking input