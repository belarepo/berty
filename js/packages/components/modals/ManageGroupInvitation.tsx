import React, { useState } from 'react'
import { View, TouchableOpacity, TextInput } from 'react-native'
import { Buffer } from 'buffer'
import { Icon } from '@ui-kitten/components'
import { CommonActions, useNavigation } from '@react-navigation/native'
import { useTranslation } from 'react-i18next'

import { useStyles } from '@berty/styles'
import messengerMethodsHooks from '@berty/store/methods'
import { dispatch as navDispatch } from '@berty/navigation/rootRef'
import { useThemeColor } from '@berty/store'
import { useAppDispatch, useConversation } from '@berty/hooks'

import { TabBar } from '../shared-components/TabBar'
import { FingerprintContent } from '../shared-components/FingerprintContent'
import InvalidScan from './InvalidScan'
import { MultiMemberAvatar } from '../avatars'
import { UnifiedText } from '../shared-components/UnifiedText'

const useStylesModal = () => {
	const [{ width, border, height, opacity }] = useStyles()
	return {
		closeRequest: [width(45), height(45), border.radius.scale(22.5)],
		closeRequestIcon: opacity(0.5),
	}
}

const BodyManageGroupInvitationContent: React.FC<{}> = ({ children }) => {
	const [{ margin }] = useStyles()
	return (
		<View style={[margin.top.big]}>
			<View>{children}</View>
		</View>
	)
}

const SelectedContent = ({
	contentName,
	pubKey,
	isEncrypted,
}: {
	contentName: string
	pubKey: string
	isEncrypted: boolean
}) => {
	const [{ padding }] = useStyles()
	switch (contentName.toLowerCase()) {
		case 'fingerprint':
			return <FingerprintContent seed={pubKey} isEncrypted={isEncrypted} />
		default:
			return (
				<UnifiedText style={[padding.horizontal.medium]}>
					Error: Unknown content name "{contentName}"
				</UnifiedText>
			)
	}
}

export const ManageGroupInvitation: React.FC<{
	link: string
	displayName: string
	publicKey: string
	type: string
	isPassword: boolean
}> = ({ link, type, displayName, publicKey, isPassword }) => {
	const navigation = useNavigation()
	const { call: joinConversation, done, error } = messengerMethodsHooks.useConversationJoin()
	const [{ row, text, column, flex, absolute, padding, border, margin }] = useStyles()
	const colors = useThemeColor()
	const [selectedContent, setSelectedContent] = useState('fingerprint')
	const _styles = useStylesModal()
	const { t }: any = useTranslation()
	const dispatch = useAppDispatch()
	const { dispatch: navigationDispatch } = useNavigation()
	const convId = useConversation(publicKey)?.publicKey

	const [password, setPassword] = useState('')

	// TODO: handle error (shouldn't happen since we checked the link previously, but still)

	React.useEffect(() => {
		if (done && !error) {
			navDispatch(
				CommonActions.reset({
					routes: [{ name: 'Main.Home' }],
				}),
			)
		}
	}, [done, error, dispatch])
	if (convId) {
		navigationDispatch(
			CommonActions.reset({
				routes: [
					{ name: 'Main.Home' },
					{
						name: 'Chat.Group',
						params: {
							convId,
						},
					},
				],
			}),
		)
	}
	if (error) {
		return <InvalidScan type={type} error={error} />
	}

	return (
		<View
			style={[{ justifyContent: 'center', alignItems: 'center', height: '100%' }, padding.medium]}
		>
			<View
				style={[
					padding.horizontal.medium,
					padding.bottom.medium,
					border.radius.large,
					{ width: '100%', backgroundColor: colors['main-background'] },
				]}
			>
				<View style={[absolute.scale({ top: -50 }), row.item.justify]}>
					<MultiMemberAvatar
						publicKey={publicKey}
						fallbackNameSeed={displayName}
						style={[border.shadow.big, row.center, { shadowColor: colors.shadow }] as any}
						size={100}
					/>
				</View>
				<View style={[padding.top.scale(55)]}>
					<UnifiedText style={{ textAlign: 'center' }}>{displayName}</UnifiedText>
					<TabBar
						tabs={[
							{
								key: 'fingerprint',
								name: t('modals.group-invitation.fingerprint'),
								icon: 'fingerprint',
								iconPack: 'custom',
							},
							{
								key: 'info',
								name: t('modals.group-invitation.info'),
								icon: 'info-outline',
								buttonDisabled: true,
							},
							{
								key: 'devices',
								name: t('modals.group-invitation.devices'),
								icon: 'smartphone',
								iconPack: 'feather',
								iconTransform: [{ rotate: '22.5deg' }, { scale: 0.8 }],
								buttonDisabled: true,
							},
						]}
						onTabChange={setSelectedContent}
					/>
					<BodyManageGroupInvitationContent>
						<SelectedContent
							contentName={selectedContent}
							pubKey={publicKey}
							isEncrypted={isPassword}
						/>
					</BodyManageGroupInvitationContent>
				</View>
				{isPassword ? (
					<View>
						<View
							style={[
								{ flexDirection: 'row', justifyContent: 'center', alignItems: 'center' },
								margin.top.medium,
							]}
						>
							<Icon name='info-outline' fill={colors['background-header']} width={15} height={15} />
							<UnifiedText
								style={[
									{
										color: colors['background-header'],
										paddingLeft: 5,
									},
									text.size.small,
									text.align.center,
									text.light,
								]}
							>
								{t('modals.group-invitation.password-label')}
							</UnifiedText>
						</View>
						<View
							style={[
								border.radius.small,
								padding.small,
								margin.top.medium,
								row.fill,
								padding.vertical.scale(12),
								{ backgroundColor: colors['negative-asset'] },
							]}
						>
							<TextInput
								value={password}
								secureTextEntry={true}
								onChangeText={setPassword}
								autoCapitalize='none'
								editable={true}
								style={[{ fontFamily: 'Open Sans' }, text.light, { width: '100%' }]}
								placeholder={t('modals.group-invitation.password-placeholder')}
							/>
						</View>
					</View>
				) : null}
				<View style={[padding.top.big, row.fill, padding.medium]}>
					<TouchableOpacity
						onPress={() => {
							isPassword
								? joinConversation({ link, passphrase: Buffer.from(password) })
								: joinConversation({ link })
						}}
						style={[
							flex.medium,
							padding.vertical.scale(12),
							border.radius.small,
							{ backgroundColor: colors['positive-asset'] },
						]}
					>
						<UnifiedText style={{ textAlign: 'center', color: colors['background-header'] }}>
							{t('modals.group-invitation.join')}
						</UnifiedText>
					</TouchableOpacity>
				</View>
			</View>
			<TouchableOpacity
				style={[
					padding.vertical.medium,
					border.shadow.medium,
					row.item.justify,
					column.justify,
					_styles.closeRequest,
					{
						position: 'absolute',
						bottom: '2%',
						backgroundColor: colors['main-background'],
						shadowColor: colors.shadow,
					},
				]}
				onPress={navigation.goBack}
			>
				<Icon
					style={[row.item.justify, _styles.closeRequestIcon]}
					name='close-outline'
					width={25}
					height={25}
					fill={colors['secondary-text']}
				/>
			</TouchableOpacity>
		</View>
	)
}
